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

package market

import (
	"fmt"

	"github.com/awishformore/m3/adaptor/contract"
	"github.com/awishformore/m3/binding"
	"github.com/awishformore/m3/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Maker uses a market contract and a simple proxy contract to execute
// trades on the given blockchain in an atomic way.
type Maker struct {
	*contract.Ethereum
	market *binding.SimpleMarket
}

// NewMaker creates a new wrapper around a couple of contracts to implement
// all required functionality on top of the blockchain.
func NewMaker(backend bind.ContractBackend, address common.Address) (*Maker, error) {

	// bind market contract
	market, err := binding.NewSimpleMarket(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind market contract: %v (%v)", address, err)
	}

	mm := Maker{
		Ethereum: contract.NewEthereum(backend, address),
		market:   market,
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
