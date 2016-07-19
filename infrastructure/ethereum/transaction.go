package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// Transaction represents a wrapper around an ethereum transaction to implement
// our simplified interface.
type Transaction struct {
	inner *types.Transaction
}

// From returns the ethereum address hash of the sender as a slice of bytes.
func (tx *Transaction) From() []byte {
	address, _ := tx.inner.From()
	return []byte(address[:])
}

// To returns the ethereum address hash of the recipient as a slice of bytes.
func (tx *Transaction) To() []byte {
	address := tx.inner.To()
	return []byte(address[:])
}

// Cost returns the cost in wei of the transaction as big int.
func (tx *Transaction) Cost() *big.Int {
	return tx.inner.Cost()
}

// Value returns the transmitted value in wei for the transaction as big int.
func (tx *Transaction) Value() *big.Int {
	return tx.inner.Value()
}

// Data returns the transmitted data of the transaction as a byte slice.
func (tx *Transaction) Data() []byte {
	return tx.inner.Data()
}

// Hash returns the transaction unique identification hash as a byte slice.
func (tx *Transaction) Hash() []byte {
	hash := tx.inner.Hash()
	return []byte(hash[:])
}
