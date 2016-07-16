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

package model

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Order represents an order on maker market.
type Order struct {
	ID         *big.Int
	SellToken  common.Address
	SellAmount *big.Int
	BuyToken   common.Address
	BuyAmount  *big.Int
}

// Rate will return the rate between sell and buy amounts (for bids).
func (o Order) Rate() *big.Rat {
	return new(big.Rat).SetFrac(o.SellAmount, o.BuyAmount)
}

// Inverse returns the inverted rate between sell and buy amounts (for asks).
func (o Order) Inverse() *big.Rat {
	return new(big.Rat).Inv(o.Rate())
}

// ToSellAmount returns a buy token amount in equivalent of sell tokens.
func (o Order) ToSellAmount(buyAmount *big.Int) *big.Int {
	sellAmount := new(big.Int).Set(buyAmount)
	sellAmount.Mul(sellAmount, o.SellAmount)
	sellAmount.Div(sellAmount, o.BuyAmount)
	return sellAmount
}

// ToBuyAmount returns a sell token amount in equivalent of buy tokens.
func (o Order) ToBuyAmount(sellAmount *big.Int) *big.Int {
	buyAmount := new(big.Int).Set(sellAmount)
	buyAmount.Mul(buyAmount, o.BuyAmount)
	buyAmount.Div(buyAmount, o.SellAmount)
	return buyAmount
}

// String returns a human readable format of the order.
func (o Order) String() string {
	return fmt.Sprintf("%v\t(%v:%v)", o.Rate(), o.SellAmount, o.BuyAmount)
}

// Valid returns false if either sell or buy amount is zero.
func (o Order) Valid() bool {
	zero := big.NewInt(0)
	if o.SellAmount.Cmp(zero) == 0 {
		return false
	}
	if o.BuyAmount.Cmp(zero) == 0 {
		return false
	}
	return true
}
