package business

import "github.com/ethereum/go-ethereum/common"

// Contract represents a common interface for ethereum contracts.
type Contract interface {
	Address() common.Address
}
