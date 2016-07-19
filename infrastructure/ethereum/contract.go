package ethereum

import "math/big"

// Contract is a wrapper around the ethereum bound contract to implement our
// simplified contract interface.
type Contract struct {
}

// Call attempts to call the given method with given input and extracting the
// result into the given output.
func Call(method string, input interface{}, output interface{}) error {
	return nil
}

// Transfer attempts to transfer the desired amount of ether to the given
// contract.
func Transfer(amount *big.Int) (*Transaction, error) {
	return nil, nil
}

// Transact tries to call the given method with given input and returns a
// transaction wrapper with the details of the transaction.
func Transact(method string, amount *big.Int, input interface{}) (*Transaction, error) {
	return nil, nil
}
