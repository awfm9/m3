package model

import "math/big"

// Order represents an order on maker market.
type Order struct {
	ID         *big.Int
	BuyAmount  *big.Int
	BuyToken   string
	SellAmount *big.Int
	SellToken  string
	Rate       *big.Rat
}
