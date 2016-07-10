package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Margin represents a margin made in a given token.
type Margin struct {
	Token  common.Address
	Amount *big.Int
}
