//go:generate abigen --sol contract.sol --pkg contract --out contract.go

// Package contract contains ethereum smart contracts written in solidity. This
// file contains a list of instructions to automatically compile them into the
// corresponding Go bindings.
package contract
