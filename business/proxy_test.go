package business

import (
	"math/big"
	"math/rand"

	"github.com/awishformore/m3/model"
)

type fakeProxy struct {
}

func (*fakeProxy) Atomic(*model.Order, *model.Order) (*model.Twin, error) {
	first := model.Margin{
		Token:  randomAddress(),
		Amount: big.NewInt(rand.Int63()),
	}
	second := model.Margin{
		Token:  randomAddress(),
		Amount: big.NewInt(rand.Int63()),
	}
	twin := model.Twin{
		First:  &first,
		Second: &second,
		Cost:   big.NewInt(rand.Int63()),
	}
	return &twin, nil
}
