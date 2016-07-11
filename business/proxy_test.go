package business

import (
	"fmt"
	"math/big"
	"math/rand"

	"github.com/awishformore/m3/model"
	"github.com/ethereum/go-ethereum/common"
)

type fakeProxy struct {
}

func (*fakeProxy) Atomic(*model.Order, *model.Order) (*model.Twin, error) {
	alpha := make([]byte, 20)
	beta := make([]byte, 20)
	_, err := rand.Read(alpha)
	if err != nil {
		return nil, fmt.Errorf("could not get random alpha token address (%v)", err)
	}
	_, err = rand.Read(beta)
	if err != nil {
		return nil, fmt.Errorf("could not get random beta token addres (%v)", err)
	}
	first := model.Margin{
		Token:  common.BytesToAddress(alpha),
		Amount: big.NewInt(rand.Int63()),
	}
	second := model.Margin{
		Token:  common.BytesToAddress(beta),
		Amount: big.NewInt(rand.Int63()),
	}
	twin := model.Twin{
		First:  &first,
		Second: &second,
		Cost:   big.NewInt(rand.Int63()),
	}
	return &twin, nil
}
