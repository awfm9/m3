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
	"time"

	"github.com/awishformore/logger"

	"github.com/awishformore/m3/adaptor"
)

// Matcher is a market matcher that will try to match overlapping orders against
// each other.
type Matcher struct {
	log      *logger.Logger
	market   adaptor.Market
	done     chan struct{}
	interval time.Duration
	books    map[string]*Book
}

// NewMatcher creates a new market matcher that will try to execute trades against each other.
func NewMatcher(log *logger.Logger, market adaptor.Market, interval time.Duration) *Matcher {

	// create the channel to signal shutdown
	m := Matcher{
		log:      log,
		market:   market,
		done:     make(chan struct{}),
		interval: interval,
	}

	return &m
}

// Stop will end the execution loop of the matcher and return after cleanly
// shutting down.
func (m *Matcher) Stop() {
	m.done <- struct{}{}
	<-m.done
}

// start will start the matcher execution loop.
func (m *Matcher) start() {

	// start the timer with the configured timeout
	ticker := time.NewTicker(m.interval)

	// execution loop of the matcher where it does its stuff
MatcherLoop:
	for {
		select {
		case <-m.done:
			break MatcherLoop

		// once every interval we get all orders from the contract and match them
		case <-ticker.C:

			// try getting all the orders from the contract
			err := m.getOrders()
			if err != nil {
				m.log.Errorf("could not get orders (%v)", err)
				continue
			}

			// try matching the orders for trade opportunities
			err = m.doArbitrage()
			if err != nil {
				m.log.Errorf("could not do arbitrage (%v)", err)
				continue
			}
		}
	}

	// clean up
	ticker.Stop()

	// close the done channel so stop returns
	close(m.done)
}

func (m *Matcher) getOrders() error {

	// clear all books
	m.books = make(map[string]*Book)

	// retrieve valid orders from contract
	orders, err := m.market.Orders()
	if err != nil {
		return fmt.Errorf("could not retrieve orders from market (%v)", err)
	}

	// put the orders into the respective order books for their pair
	for _, order := range orders {

		// check for both pair as bid and pair as ask
		bidPair := order.BuyToken + order.SellToken
		askPair := order.SellToken + order.BuyToken

		// check if there is a book with the bid pair and add as bid if found
		bidBook, ok := m.books[bidPair]
		if ok {
			bidBook.AddBid(order)
			continue
		}

		// check if there is a book with the ask pair and add as ask if found
		askBook, ok := m.books[askPair]
		if ok {
			askBook.AddAsk(order)
			continue
		}

		// if no book was found for pair or inversed pair, create bid book
		book := Book{}
		book.AddBid(order)
		m.books[bidPair] = &book
	}

	return nil
}

func (m *Matcher) doArbitrage() error {

	// for each book, check if there are overlapping orders
	for _, book := range m.books {

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
			if bid.Rate.Cmp(ask.Rate) <= 0 {
				break
			}

			// execute the trade
			m.market.Trade(bid.ID, ask.ID)
			if err != nil {
				m.log.Warningf("failed to execute trade: %v & %v", bid.ID, ask.ID)
				break
			}
		}
	}

	return nil
}
