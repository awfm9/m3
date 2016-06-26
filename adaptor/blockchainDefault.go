package adaptor

import (
	"fmt"

	"github.com/awishformore/m3/contract"
	"github.com/awishformore/m3/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// BlockchainDefault implements a blockchain interface using a market contract
// and a proxy contract for trading atomically.
type BlockchainDefault struct {
	conn    rpc.Client
	backend bind.ContractBackend
	market  *contract.SimpleMarket
	proxy   *contract.TradeProxy
}

// NewBlockchain creates a new default blockchain wrapper.
func NewBlockchain(ipc string, marketAddress string, proxyAddress string) (*BlockchainDefault, error) {

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

	bc := BlockchainDefault{
		conn:    conn,
		backend: backend,
		market:  market,
		proxy:   proxy,
	}

	return &bc, nil
}

// Close will free up the resources associated with the blockchain wrapper.
func (bc *BlockchainDefault) Close() {
	bc.conn.Close()
}

// GetOrders will get a slice of all orders on the market given.
func (bc *BlockchainDefault) GetOrders() ([]*model.Order, error) {

	// slice to hold valid orders
	orders := []*model.Order{}

	// get the last order ID to iterate through all existing orders
	id, err := bc.market.Last_offer_id(nil)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve last offer ID (%v)", err)
	}

	// while we have a positive ID
	for id.Sign() > 0 {

		// get the info for this order
		info, err := bc.market.Offers(nil, id)
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
