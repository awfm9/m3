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

package wallet

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/awishformore/m3/adaptor/contract"
	"github.com/awishformore/m3/binding"
	"github.com/awishformore/m3/business"
	"github.com/awishformore/m3/model"
)

// M3Wallet represents a custom wallet used in this daemon, capable of executing
// atomic trades against a market.
type M3Wallet struct {
	*contract.Ethereum
	wallet *binding.M3Wallet
}

// NewM3 creates a new M3 wallet.
func NewM3(backend bind.ContractBackend, address common.Address) (*M3Wallet, error) {

	// bind market contract
	wallet, err := binding.NewM3Wallet(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind market contract: %v (%v)", address, err)
	}

	m3 := M3Wallet{
		Ethereum: contract.NewEthereum(backend, address),
		wallet:   wallet,
	}

	return &m3, nil
}

// Balance returns the balance currently held in our proxy.
func (mw *M3Wallet) Balance(address common.Address) (*big.Int, error) {
	balance, err := mw.wallet.BalanceOf(nil, address)
	if err != nil {
		return nil, fmt.Errorf("could not get balance from proxy (%v)", err)
	}

	return balance, nil
}

// ExecuteAtomic will execute the trades with the two given orders.
func (mw *M3Wallet) ExecuteAtomic(market business.Market, first *model.Order, firstSelling *big.Int, second *model.Order, secondSelling *big.Int) (*big.Int, error) {
	info, err := mw.wallet.AtomicTradePair(nil, market.Address(), first.ID, firstSelling, second.ID, secondSelling)
	if err != nil {
		return nil, fmt.Errorf("could not atomically execute trades (%v)", err)
	}
	return info.Cost(), nil
}
