// Copyright (c) 2015 Max Wolter
//
// This file is part of M3 - Maker Market Maker.
//
// M3 is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// M3 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with M3.  If not, see <http://www.gnu.org/licenses/>.

package atomic

import (
	"fmt"
	"math/big"

	"github.com/awishformore/m3/contract"
	"github.com/awishformore/m3/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Maker uses a market contract and a simple proxy contract to execute
// trades on the given blockchain in an atomic way.
type Maker struct {
	backend bind.ContractBackend
	address common.Address
	market  *contract.SimpleMarket
	proxy   *contract.TraderKeeper
}

// NewMaker creates a new wrapper around a couple of contracts to implement
// all required functionality on top of the blockchain.
func NewMaker(backend bind.ContractBackend, maker common.Address, keeper common.Address) (*Maker, error) {

	// bind market contract
	market, err := contract.NewSimpleMarket(maker, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind market contract: %v (%v)", market, err)
	}

	// bind proxy contract
	proxy, err := contract.NewTraderKeeper(keeper, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind proxy contract: %v (%v)", proxy, err)
	}

	mm := Maker{
		backend: backend,
		address: maker,
		market:  market,
		proxy:   proxy,
	}

	return &mm, nil
}

// Orders will return a slice of all active orders on the market.
func (mm *Maker) Orders() ([]*model.Order, error) {

	// slice to hold valid orders
	orders := []*model.Order{}

	// get the last order ID to iterate through all existing orders
	id, err := mm.market.Last_offer_id(nil)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve last offer ID (%v)", err)
	}

	// while we have a positive ID
	for id.Sign() > 0 {

		// get the info for this order
		info, err := mm.market.Offers(nil, id)
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
			BuyAmount:  info.Buy_how_much,
			BuyToken:   info.Buy_which_token,
			SellAmount: info.Sell_how_much,
			SellToken:  info.Sell_which_token,
		}

		// append the order
		orders = append(orders, &order)
	}

	return orders, nil
}

// ExecuteAtomic will execute the trades with the two given orders.
func (mm *Maker) ExecuteAtomic(first *model.Order, second *model.Order) (*model.Twin, error) {

	// check our balances before the trade
	beforeFirst, err := mm.proxy.BalanceOf(nil, first.BuyToken)
	if err != nil {
		return nil, fmt.Errorf("could not check first token before twin (%v)", err)
	}
	beforeSecond, err := mm.proxy.BalanceOf(nil, first.SellToken)
	if err != nil {
		return nil, fmt.Errorf("could not check second token before twin (%v)", err)
	}

	// execute the trades using the market proxy
	info, err := mm.proxy.Trade(nil, first.ID, second.ID, first.BuyToken, second.BuyToken, mm.address)
	if err != nil {
		return nil, fmt.Errorf("could not atomically execute trades (%v)", err)
	}

	// check our balances after the trade
	afterFirst, err := mm.proxy.BalanceOf(nil, first.BuyToken)
	if err != nil {
		return nil, fmt.Errorf("could not check first token after twin (%v)", err)
	}
	afterSecond, err := mm.proxy.BalanceOf(nil, first.SellToken)
	if err != nil {
		return nil, fmt.Errorf("could not check second token after twin (%v)", err)
	}

	// calculate margins
	marginFirst := model.Margin{
		Amount: new(big.Int).Sub(beforeFirst, afterFirst),
		Token:  first.BuyToken,
	}
	marginSecond := model.Margin{
		Amount: new(big.Int).Sub(beforeSecond, afterSecond),
		Token:  first.SellToken,
	}

	// create the twin
	twin := model.Twin{
		First:  &marginFirst,
		Second: &marginSecond,
		Cost:   info.Cost(),
	}

	return &twin, nil
}

// Balance returns the balance currently held in our proxy.
func (mm *Maker) Balance(address common.Address) (*big.Int, error) {

	// try to get the balance
	balance, err := mm.proxy.BalanceOf(nil, address)
	if err != nil {
		return nil, fmt.Errorf("could not get balance from proxy (%v)", err)
	}

	return balance, nil
}
