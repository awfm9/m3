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
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func randomAddress() common.Address {
	bytes := make([]byte, 20)
	_, _ = rand.Read(bytes)
	return common.BytesToAddress(bytes)
}

func randomDuration() time.Duration {
	seconds := rand.Uint32()
	return time.Duration(seconds) * time.Second
}

func TestNewMatcher(t *testing.T) {
	log := &fakeLog{}
	market := &fakeMarket{}
	proxy := &fakeProxy{}
	wallet := &fakeWallet{}
	threshold := uint64(rand.Uint32())
	refresh := randomDuration()
	matcher := NewMatcher(
		log, market, proxy, wallet,
		SetThreshold(threshold),
		SetRefresh(refresh),
	)
	if matcher.log != log {
		t.Errorf("did not save log reference on creation")
	}
	if matcher.market != market {
		t.Errorf("did not save market reference on creation")
	}
	if matcher.proxy != proxy {
		t.Errorf("did not save proxy referenc on creation")
	}
	if matcher.wallet != wallet {
		t.Errorf("did not save wallet reference on creation")
	}
	if matcher.threshold.Cmp(new(big.Int).SetUint64(threshold)) != 0 {
		t.Errorf("did not save threshold amount on creation")
	}
	if matcher.refresh != refresh {
		t.Errorf("did not save refresh duration on creation")
	}
	if matcher.done == nil {
		t.Errorf("did not create done channel on creation")
	}
}
