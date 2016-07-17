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

	"github.com/awishformore/m3/model"
	"github.com/ethereum/go-ethereum/common"
)

// Wallet wraps a wallet address with all necessary functions to transfer ERC20
// tokens and grant rights to transfer them on our behalf.
type Wallet interface {
	Contract
	Balance(token common.Address) (*big.Int, error)
	ExecuteAtomic(market Market, first *model.Order, firstSelling *big.Int, second *model.Order, secondSelling *big.Int) (*big.Int, error)
}
