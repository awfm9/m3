package business

import (
	"math/big"
	"math/rand"
	"reflect"
	"testing"

	"github.com/awishformore/m3/model"
)

func TestAddHighestBid(t *testing.T) {

	// create some test orders
	o1 := &model.Order{
		ID:         big.NewInt(rand.Int63()),
		BuyAmount:  big.NewInt(100),
		SellAmount: big.NewInt(110),
	}
	o2 := &model.Order{
		ID:         big.NewInt(rand.Int63()),
		BuyAmount:  big.NewInt(100),
		SellAmount: big.NewInt(100),
	}
	o3 := &model.Order{
		ID:         big.NewInt(rand.Int63()),
		BuyAmount:  big.NewInt(100),
		SellAmount: big.NewInt(90),
	}

	// create test vectors
	vectors := map[string]struct {
		orders  []*model.Order
		bids    []*model.Order
		fail    bool
		highest *model.Order
		sorted  []*model.Order
	}{
		"none": {
			orders:  []*model.Order{},
			bids:    nil,
			fail:    true,
			highest: nil,
			sorted:  nil,
		},
		"single": {
			orders:  []*model.Order{o1},
			bids:    []*model.Order{o1},
			fail:    false,
			highest: o1,
			sorted:  []*model.Order{o1},
		},
		"first": {
			orders:  []*model.Order{o1, o2, o3},
			bids:    []*model.Order{o1, o2, o3},
			fail:    false,
			highest: o1,
			sorted:  []*model.Order{o1, o2, o3},
		},
		"second": {
			orders:  []*model.Order{o2, o1, o3},
			bids:    []*model.Order{o2, o1, o3},
			fail:    false,
			highest: o1,
			sorted:  []*model.Order{o1, o2, o3},
		},
		"third": {
			orders:  []*model.Order{o3, o2, o1},
			bids:    []*model.Order{o3, o2, o1},
			fail:    false,
			highest: o1,
			sorted:  []*model.Order{o1, o2, o3},
		},
	}

	// go through test vectors and check if all bids are added correctly
	for name, vector := range vectors {

		// create the book
		book := &Book{}

		// add the orders one by one
		for _, order := range vector.orders {
			book.AddBid(order)
		}

		// check if the bids correspond to the orders
		if !reflect.DeepEqual(book.bids, vector.bids) {
			t.Errorf("[%v] bids in book do not match expected:\n%v\n%v", name, book.bids, vector.bids)
		}

		// try to get the highest bid
		bid, err := book.HighestBid()
		if (err != nil) != vector.fail {
			t.Errorf("[%v] fail condition on case wrong: %v != %v", name, err != nil, vector.fail)
		}

		// skip those that failed
		if err != nil {
			continue
		}

		// check that the highest bid is as expected
		if !reflect.DeepEqual(bid, vector.highest) {
			t.Errorf("[%v] highest bid not as expected:\n%v\n%v", name, bid, vector.highest)
		}

		// check that all bids are ordered correctly now
		if !reflect.DeepEqual(book.bids, vector.sorted) {
			t.Errorf("[%v] sorted bids not as expected:\n%v\n%v", name, book.bids, vector.sorted)
		}
	}
}

func TestAddLowestAsk(t *testing.T) {

	// create some test orders
	o1 := &model.Order{
		ID:         big.NewInt(rand.Int63()),
		BuyAmount:  big.NewInt(100),
		SellAmount: big.NewInt(110),
	}
	o2 := &model.Order{
		ID:         big.NewInt(rand.Int63()),
		BuyAmount:  big.NewInt(100),
		SellAmount: big.NewInt(100),
	}
	o3 := &model.Order{
		ID:         big.NewInt(rand.Int63()),
		BuyAmount:  big.NewInt(100),
		SellAmount: big.NewInt(90),
	}

	// create test vectors
	vectors := map[string]struct {
		orders []*model.Order
		asks   []*model.Order
		fail   bool
		lowest *model.Order
		sorted []*model.Order
	}{
		"none": {
			orders: []*model.Order{},
			asks:   nil,
			fail:   true,
			lowest: nil,
			sorted: nil,
		},
		"single": {
			orders: []*model.Order{o1},
			asks:   []*model.Order{o1},
			fail:   false,
			lowest: o1,
			sorted: []*model.Order{o1},
		},
		"first": {
			orders: []*model.Order{o1, o2, o3},
			asks:   []*model.Order{o1, o2, o3},
			fail:   false,
			lowest: o3,
			sorted: []*model.Order{o3, o2, o1},
		},
		"second": {
			orders: []*model.Order{o2, o1, o3},
			asks:   []*model.Order{o2, o1, o3},
			fail:   false,
			lowest: o3,
			sorted: []*model.Order{o3, o2, o1},
		},
		"third": {
			orders: []*model.Order{o3, o2, o1},
			asks:   []*model.Order{o3, o2, o1},
			fail:   false,
			lowest: o3,
			sorted: []*model.Order{o3, o2, o1},
		},
	}

	// go through test vectors and check if all asks are added correctly
	for name, vector := range vectors {

		// create the book
		book := &Book{}

		// add the orders one by one
		for _, order := range vector.orders {
			book.AddAsk(order)
		}

		// check if the asks correspond to the orders
		if !reflect.DeepEqual(book.asks, vector.asks) {
			t.Errorf("[%v] asks in book do not match expected:\n%v\n%v", name, book.asks, vector.asks)
		}

		// try to get the lowest ask
		ask, err := book.LowestAsk()
		if (err != nil) != vector.fail {
			t.Errorf("[%v] fail condition on case wrong: %v != %v", name, err != nil, vector.fail)
		}

		// skip those that failed
		if err != nil {
			continue
		}

		// check that the lowest ask is as expected
		if !reflect.DeepEqual(ask, vector.lowest) {
			t.Errorf("[%v] lowest ask not as expected:\n%v\n%v", name, ask, vector.lowest)
		}

		// check that all asks are ordered correctly now
		if !reflect.DeepEqual(book.asks, vector.sorted) {
			t.Errorf("[%v] sorted asks not as expected:\n%v\n%v", name, book.asks, vector.sorted)
		}
	}
}
