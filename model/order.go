package model

import "math/big"

// Order represents an order on maker market.
type Order struct {
	ID         *big.Int
	SellAmount *big.Int
	SellToken  string
	BuyAmount  *big.Int
	BuyToken   string
}
