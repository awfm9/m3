package adaptor

import (
	"math/big"

	"github.com/awishformore/m3/model"
)

// Market is an interface to interact with a market located on the blockchain.
type Market interface {
	Orders() ([]*model.Order, error)
	Trade(*big.Int, *big.Int) error
	Close()
}
