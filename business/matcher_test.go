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
	"encoding/binary"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func randAddress() common.Address {
	bytes := make([]byte, 20)
	_, _ = rand.Read(bytes)
	return common.BytesToAddress(bytes)
}

func randDuration() time.Duration {
	seconds := rand.Uint32()
	return time.Duration(seconds) * time.Second
}

func randUint64() uint64 {
	bytes := make([]byte, 8)
	_, _ = rand.Read(bytes)
	return binary.BigEndian.Uint64(bytes)
}

func randInt64() int64 {
	bytes := make([]byte, 8)
	_, _ = rand.Read(bytes)
	return rand.Int63()
}

func TestNewMatcher(t *testing.T) {
	log := &fakeLogger{}
	atomic := &fakeAtomic{}
	threshold := randUint64()
	refresh := randDuration()
	matcher := NewMatcher(
		log, atomic,
		SetThreshold(threshold),
		SetRefresh(refresh),
	)
	if matcher.log != log {
		t.Errorf("did not save log reference on creation")
	}
	if matcher.atomic != atomic {
		t.Errorf("did not save atomic reference on creation")
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
