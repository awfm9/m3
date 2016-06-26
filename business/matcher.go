package business

import (
	"time"

	"github.com/awishformore/m3/contract"
)

// Matcher is a market matcher that will try to match overlapping orders against
// each other.
type Matcher struct {
}

// NewMatcher creates a new market matcher that will try to execute trades against each other.
func NewMatcher(market *contract.SimpleMarket, proxy *contract.TradeProxy) (*Matcher, error) {
	return &Matcher{}, nil
}

// Wait will wait for the matcher to finish.
func (m *Matcher) Wait() {
	time.Sleep(time.Second)
}
