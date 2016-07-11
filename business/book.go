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
	"fmt"
	"sort"

	"github.com/awishformore/m3/model"
	"github.com/ethereum/go-ethereum/common"
)

// Book represents an order book for either bids or asks.
type Book struct {
	Base  common.Address
	Quote common.Address
	bids  []*model.Order
	asks  []*model.Order
}

// AddBid will add an order to the orders.
func (b *Book) AddBid(order *model.Order) {
	b.bids = append(b.bids, order)
}

// HighestBid will return the highest bid order.
func (b *Book) HighestBid() (*model.Order, error) {
	if len(b.bids) == 0 {
		return nil, fmt.Errorf("no bids in book")
	}
	sort.Sort(ByRateDesc(b.bids))
	return b.bids[0], nil
}

// AddAsk will sort the orders by increasing price.
func (b *Book) AddAsk(order *model.Order) {
	b.asks = append(b.asks, order)
}

// LowestAsk will return the lowest ask order.
func (b *Book) LowestAsk() (*model.Order, error) {
	if len(b.asks) == 0 {
		return nil, fmt.Errorf("no asks in book")
	}
	sort.Sort(ByRateAsc(b.asks))
	return b.asks[0], nil
}

// ByRateDesc defines sorting by descending rate, putting the top rate as the first element.
type ByRateDesc []*model.Order

func (brd ByRateDesc) Len() int               { return len(brd) }
func (brd ByRateDesc) Swap(i int, j int)      { brd[i], brd[j] = brd[j], brd[i] }
func (brd ByRateDesc) Less(i int, j int) bool { return brd[i].Rate().Cmp(brd[j].Rate()) > 0 }

// ByRateAsc defines sorting by ascending rate, putting the bottom rate as the first element.
type ByRateAsc []*model.Order

func (brd ByRateAsc) Len() int               { return len(brd) }
func (brd ByRateAsc) Swap(i int, j int)      { brd[i], brd[j] = brd[j], brd[i] }
func (brd ByRateAsc) Less(i int, j int) bool { return brd[i].Rate().Cmp(brd[j].Rate()) < 0 }
