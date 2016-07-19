package adaptor

import "math/big"

// Contract interface is implemented by contracts on different networks.
type Contract interface {
	Call(method string, input interface{}, output interface{})
	Transfer(amount *big.Int) (Transaction, error)
	Transact(method string, amount *big.Int, input interface{}) (Transaction, error)
}
