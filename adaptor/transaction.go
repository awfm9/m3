package adaptor

import "math/big"

// Transaction interface is implemented by transactions on different networks.
type Transaction interface {
	From() []byte
	To() []byte
	Cost() *big.Int
	Value() *big.Int
	Data() []byte
	Hash() []byte
}
