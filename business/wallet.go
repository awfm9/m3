package business

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Wallet represents a wallet or contract with a balance for various tokens.
type Wallet interface {
	Balance(common.Address) (*big.Int, error)
}
