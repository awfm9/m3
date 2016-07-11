package business

import (
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
)

func randomAddress() common.Address {
	bytes := make([]byte, 20)
	_, _ = rand.Read(bytes)
	return common.BytesToAddress(bytes)
}
