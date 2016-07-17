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
)

// Matcher is a market matcher that will try to match overlapping orders against
// each other.
type Matcher struct {
	log       Logger
	market    Market
	wallet    Wallet
	threshold *big.Int
	refresh   time.Duration
	done      chan struct{}
}

// NewMatcher creates a new market matcher that will try to execute trades against each other.
func NewMatcher(log Logger, market Market, wallet Wallet, options ...func(*Matcher)) *Matcher {

	// create the channel to signal shutdown
	m := Matcher{
		log:       log,
		market:    market,
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
			m.arbitrage(books)
			if err != nil {
				m.log.Errorf("could not compute arbitrage orders (%v)", err)
				continue
			}
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
	orders, err := market.Orders()
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
		book := Book{}
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

func (m *Matcher) arbitrage(books []*Book) {

	// for each book, check if there are overlapping orders
	for pair, book := range books {

		m.log.Infof("trading book for pair: %v (%v bids, %v asks)", pair, len(book.bids), len(book.asks))

		// keep processing book until no matching orders
		for {

			// get highest bid, go to next book if none found
			bid, err := book.HighestBid()
			if err != nil {
				m.log.Infof("could not retrieve highest bid (%v)", err)
				break
			}

			// get lowest ask, go to next book if none found
			ask, err := book.LowestAsk()
			if err != nil {
				m.log.Infof("could not retrieve highest ask (%v)", err)
				break
			}

			// check if the prices overlap
			if bid.Rate().Cmp(ask.Inverse()) <= 0 {
				m.log.Infof("could not find overlapping orders (bid: %v, ask: %v)", bid.Rate(), ask.Inverse())
				break
			}

			// get available base amount (input to first, output of second)
			baseAvailable, err := m.wallet.Balance(bid.BuyToken)
			if err != nil {
				m.log.Errorf("could not get base balance: %v (%v)", bid.BuyToken, err)
				break
			}

			// get available quote amount (output of first, input to second)
			quoteAvailable, err := m.wallet.Balance(bid.SellToken)
			if err != nil {
				m.log.Errorf("could not get quote balance: %v (%v)", bid.SellToken, err)
				break
			}

			// if both available balances are zero, we can't trade
			zero := big.NewInt(0)
			if baseAvailable.Cmp(zero) == 0 && quoteAvailable.Cmp(zero) == 0 {
				m.log.Noticef("could not find tradable balances: %v & %v", bid.BuyToken, bid.SellToken)
				break
			}

			// calculate trade volume if trading bid first and ask second
			firstBase := Max(baseAvailable, bid.BuyAmount, ask.SellAmount)
			quoteAfter := new(big.Int).Add(quoteAvailable, bid.ToSellAmount(firstBase))
			firstQuote := Max(quoteAfter, ask.BuyAmount, bid.SellAmount)
			firstTotal := new(big.Int).Add(firstBase, ask.ToSellAmount(firstQuote))

			// calculate tradable volume if trading ask first and bid second
			secondQuote := Max(quoteAvailable, ask.BuyAmount, bid.SellAmount)
			baseAfter := new(big.Int).Add(baseAvailable, ask.ToSellAmount(secondQuote))
			secondBase := Max(baseAfter, bid.BuyAmount, ask.SellAmount)
			secondTotal := new(big.Int).Add(secondBase, ask.ToSellAmount(secondQuote))

			// check which ordering of trades moves the most volume and is thus the
			// most profitable
			if firstTotal.Cmp(secondTotal) >= 0 {
				err = m.TradePair(bid, firstBase, ask, firstQuote)
				if err != nil {
					m.log.Errorf("could not execute first pair trade (%v)", err)
					break
				}
			} else {
				err = m.TradePair(ask, secondQuote, bid, secondBase)
				if err != nil {
					m.log.Errorf("could not execute first pair trade (%v)", err)
					break
				}
			}

			// remove any order that is no longer valid (any balance zero)
			if !bid.Valid() {
				err = book.PopBid()
				if err != nil {
					m.log.Errorf("could not pop bid from book (%v)", err)
					break
				}
			}
			if !ask.Valid() {
				err = book.PopAsk()
				if err != nil {
					m.log.Errorf("could not pop ask from book (%v)", err)
					break
				}
			}
		}
	}
}

// TradePair tries to execute a pair of trades in atomic fashion and updating
// the resulting state.
func (m *Matcher) TradePair(first *model.Order, firstSell *big.Int, second *model.Order, secondSell *big.Int) error {

	// try executing the trade in atomic fashion on the blockchain
	cost, err := m.wallet.ExecuteAtomic(m.market, first, firstSell, second, secondSell)
	if err != nil {
		return fmt.Errorf("could not execute trade pair atomically (%v)", err)
	}

	// calculate the remaining traded amounts for both orders
	firstBuy := first.ToBuyAmount(firstSell)
	secondBuy := second.ToBuyAmount(secondSell)

	// calculate the differences to see our profit
	deltaFirst := new(big.Int).Sub(firstBuy, secondSell)
	deltaSecond := new(big.Int).Sub(secondBuy, firstSell)

	m.log.Infof("trade pair made %v (%v) and %v (%v) for %v (wei)",
		deltaFirst, first.SellToken, deltaSecond, second.SellToken, cost)

	// adjust the bid amounts
	first.BuyAmount.Sub(first.BuyAmount, firstBuy)
	first.SellAmount.Sub(first.SellAmount, firstSell)

	// adjust the ask amounts
	second.BuyAmount.Sub(second.BuyAmount, secondBuy)
	second.SellAmount.Sub(second.SellAmount, secondSell)

	return nil
}

// Max returns the maximum of three big ints.
func Max(numbers ...*big.Int) *big.Int {
	max := big.NewInt(0)
	for _, number := range numbers {
		if number.Cmp(max) > 0 {
			max.Set(number)
		}
	}
	return max
}
