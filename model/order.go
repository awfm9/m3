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

// Rate will return the rate between buy and sell amounts.
func (o Order) Rate() *big.Rat {
	return new(big.Rat).SetFrac(o.SellAmount, o.BuyAmount)
}

// String returns a human readable format of the order.
func (o Order) String() string {
	return fmt.Sprintf("%v\t(%v:%v)", o.Rate(), o.SellAmount, o.BuyAmount)
}
