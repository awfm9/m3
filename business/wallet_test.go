package business

import (
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
)

type fakeWallet struct {
}

func (*fakeWallet) Balance(common.Address) (*big.Int, error) {
	amount := big.NewInt(rand.Int63())
	return amount, nil
}
