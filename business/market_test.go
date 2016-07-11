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
