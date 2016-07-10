// Copyright (c) 2015 Max Wolter
//
// This file is part of M3 - Maker Market Maker.
//
// M3 is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// M3 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with M3.  If not, see <http://www.gnu.org/licenses/>.

package business

import (
	"fmt"
	"math/big"
	"time"

	"github.com/awishformore/m3/model"
	"github.com/ethereum/go-ethereum/common"
)

// Matcher is a market matcher that will try to match overlapping orders against
// each other.
type Matcher struct {
	log       Logger
	market    Market
	proxy     Proxy
	wallet    Wallet
	threshold *big.Int
	refresh   time.Duration
	done      chan struct{}
}

// NewMatcher creates a new market matcher that will try to execute trades against each other.
func NewMatcher(log Logger, market Market, proxy Proxy, wallet Wallet, options ...func(*Matcher)) *Matcher {

	// create the channel to signal shutdown
	m := Matcher{
		log:       log,
		market:    market,
		proxy:     proxy,
		wallet:    wallet,
		threshold: big.NewInt(30000),
		refresh:   time.Minute,
		done:      make(chan struct{}),
	}

	// apply the optional parameters
	for _, option := range options {
		option(&m)
	}

	// start the execution loop
	go m.start()

	return &m
}

// SetRefresh allows specifying a custom refresh interval for orders.
func SetRefresh(refresh time.Duration) func(*Matcher) {
	return func(m *Matcher) {
		m.refresh = refresh
	}
}

// SetThreshold allows specifying a minimum profit margin to make sure we don't
// spend more on fees than we get in returns.
func SetThreshold(threshold uint64) func(*Matcher) {
	return func(m *Matcher) {
		m.threshold.SetUint64(threshold)
	}
}

// start will begin the execution loop of the matcher.
func (m *Matcher) start() {

	// initialize tickers
	ticker := time.NewTicker(m.refresh)

	// run the execution loop until it quits, with all channels providing input
	// and output as parameters for easy testing
	m.run(m.done, ticker.C)

	// close channels and clean up
	ticker.Stop()
	close(m.done)
}

// Stop will end the execution loop of the matcher and return after cleanly
// shutting down.
func (m *Matcher) Stop() {
	m.done <- struct{}{}
	<-m.done
}

// start will start the matcher execution loop.
func (m *Matcher) run(done <-chan struct{}, refresh <-chan time.Time) {
Loop:
	for {
		select {

		// we received the stop signal, so quit the execution loop
		case <-done:
			break Loop

			// on every refresh interval, get all orders and try to find arbitrage
		case <-refresh:

			// try getting all the orders from the contract
			books, err := m.getBooks(m.market)
			if err != nil {
				m.log.Errorf("could not get orders (%v)", err)
				continue
			}

			// try matching the orders of each for trade opportunities
			// we need to take into account available balances and deduce them as they
			// get used up
			twins, err := m.arbitrage(books)
			if err != nil {
				m.log.Errorf("could not compute arbitrage orders (%v)", err)
				continue
			}

			// calculate total cost and margin made for each token
			cost := new(big.Int)
			changes := make(map[common.Address]*big.Int)
			for _, twin := range twins {

				// add cost
				cost.Add(cost, twin.Cost)

				// check if change in token exists, if not create, then add
				first, ok := changes[twin.First.Token]
				if !ok {
					first = big.NewInt(0)
					changes[twin.First.Token] = first
				}
				first.Add(first, twin.First.Amount)

				// check if change in second token exists, if not create and add
				second, ok := changes[twin.Second.Token]
				if !ok {
					second = big.NewInt(0)
					changes[twin.Second.Token] = second
				}
				second.Add(second, twin.Second.Amount)
			}

			m.log.Infof("executed %v twins for cost of %v wei", len(twins))

			for address, change := range changes {

				// get current balance
				balance, err := m.wallet.Balance(address)
				if err != nil {
					m.log.Warningf("could not get current balance: %v (%v)", address, err)
					continue
				}

				// print the token details and change in balance
				m.log.Infof("%v: %v (%v)", address, balance, change)
			}

			m.log.Infof("total cost: %v wei", cost)
		}
	}
}

// getBooks returs all active orders on the given maker market in the form of
// books that contain bids and asks. Each book represents one token pair, thus
// granting the application support for multiple pairs.
func (m *Matcher) getBooks(market Market) ([]*Book, error) {

	// prepare empty map with books
	bookSet := make(map[string]*Book)

	// retrieve valid orders from contract
	orders, err := m.market.Orders()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve orders from market (%v)", err)
	}

	// put the orders into the respective order bookSet for their pair
	for _, order := range orders {

		// check for both pair as bid and pair as ask
		bidPair := order.BuyToken.Hex() + order.SellToken.Hex()
		askPair := order.SellToken.Hex() + order.BuyToken.Hex()

		// check if there is a book with the bid pair and add as bid if found
		bidBook, ok := bookSet[bidPair]
		if ok {
			bidBook.AddBid(order)
			continue
		}

		// check if there is a book with the ask pair and add as ask if found
		askBook, ok := bookSet[askPair]
		if ok {
			askBook.AddAsk(order)
			continue
		}

		// if no book was found for pair or inversed pair, create bid book
		book := Book{
			Base:  order.BuyToken,
			Quote: order.SellToken,
		}
		book.AddBid(order)
		bookSet[bidPair] = &book
	}

	// turn the map into a slice
	books := make([]*Book, 0, len(bookSet))
	for _, book := range bookSet {
		books = append(books, book)
	}

	return books, nil
}

func (m *Matcher) arbitrage(books []*Book) ([]*model.Twin, error) {

	// create empty executed trades book
	twins := []*model.Twin{}

	// for each book, check if there are overlapping orders
	for _, book := range books {

		// keep processing book until no matching orders
		for {

			// get highest bid, go to next book if none found
			bid, err := book.HighestBid()
			if err != nil {
				break
			}

			// get lowest ask, go to next book if none found
			ask, err := book.LowestAsk()
			if err != nil {
				break
			}

			// check if the prices overlap
			if bid.Rate().Cmp(ask.Rate()) <= 0 {
				break
			}

			// execute the trade
			twin, err := m.proxy.Atomic(bid, ask)
			if err != nil {
				m.log.Warningf("failed to execute atomic trades: %v & %v", bid.ID, ask.ID)
				break
			}

			twins = append(twins, twin)
		}
	}

	return twins, nil
}
