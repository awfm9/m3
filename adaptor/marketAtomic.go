package adaptor

import (
	"fmt"
	"math/big"

	"github.com/awishformore/m3/contract"
	"github.com/awishformore/m3/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// AtomicMarket uses a market contract and a simple proxy contract to execute
// trades on the given blockchain in an atomic way.
type AtomicMarket struct {
	conn    rpc.Client
	backend bind.ContractBackend
	address common.Address
	market  *contract.SimpleMarket
	proxy   *contract.TradeProxy
}

// NewAtomicMarket creates a new default blockchain wrapper.
func NewAtomicMarket(ipc string, marketAddress string, proxyAddress string) (*AtomicMarket, error) {

	// initialize connection to the geth node
	conn, err := rpc.NewIPCClient(ipc)
	if err != nil {
		return nil, fmt.Errorf("could not initialize IPC connection: %v (%v)", ipc, err)
	}
	backend := backends.NewRPCBackend(conn)

	// bind maker market contract
	market, err := contract.NewSimpleMarket(common.HexToAddress(marketAddress), backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind simple market contract: %v (%v)", marketAddress, err)
	}

	// bind trade proxy contract
	proxy, err := contract.NewTradeProxy(common.HexToAddress(proxyAddress), backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind trade proxy contract: %v (%v)", proxyAddress, err)
	}

	bc := AtomicMarket{
		conn:    conn,
		backend: backend,
		address: common.HexToAddress(marketAddress),
		market:  market,
		proxy:   proxy,
	}

	return &bc, nil
}

// Close will free up the resources associated with the blockchain wrapper.
func (am *AtomicMarket) Close() {
	am.conn.Close()
}

// Orders will get a slice of all orders on the market given.
func (am *AtomicMarket) Orders() ([]*model.Order, error) {

	// slice to hold valid orders
	orders := []*model.Order{}

	// get the last order ID to iterate through all existing orders
	id, err := am.market.Last_offer_id(nil)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve last offer ID (%v)", err)
	}

	// while we have a positive ID
	for id.Sign() > 0 {

		// get the info for this order
		info, err := am.market.Offers(nil, id)
		if err != nil {
			return nil, fmt.Errorf("could not get order info: %v (%v)", id, err)
		}

		// if the order is no longer active, skip
		if !info.Active {
			continue
		}

		// build an order struct from the info
		order := model.Order{
			ID:         id,
			SellAmount: info.Sell_how_much,
			SellToken:  info.Sell_which_token.Hex(),
			BuyAmount:  info.Buy_how_much,
			BuyToken:   info.Buy_which_token.Hex(),
		}

		// append the order
		orders = append(orders, &order)
	}

	return orders, nil
}

// Trade will execute the trades with the two given orders.
func (am *AtomicMarket) Trade(firstID *big.Int, secondID *big.Int) error {

	// execute the trades using the market proxy
	_, err := am.proxy.Trade(nil, am.address, firstID, secondID)
	if err != nil {
		return fmt.Errorf("could not atomically execute trades (%v)", err)
	}

	return nil
}
