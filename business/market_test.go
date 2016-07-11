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

package business

import (
	"math/big"
	"math/rand"

	"github.com/awishformore/m3/model"
)

type fakeMarket struct {
}

func (*fakeMarket) Orders() ([]*model.Order, error) {
	n := rand.Int()%240 + 16
	orders := make([]*model.Order, 0, n)
	for i := 0; i < n; i++ {
		order := model.Order{
			ID:         big.NewInt(rand.Int63()),
			SellToken:  randomAddress(),
			SellAmount: big.NewInt(rand.Int63()),
			BuyToken:   randomAddress(),
			BuyAmount:  big.NewInt(rand.Int63()),
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
