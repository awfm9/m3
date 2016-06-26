package adaptor

import "github.com/awishformore/m3/model"

// Blockchain is used as the interface to interact with the Ethereum blockchain
// through smart contracts.
type Blockchain interface {
	GetOrders() ([]*model.Order, error)
	Close()
}
