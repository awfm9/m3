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

package main

import (
	"fmt"
	"os"
	"os/signal"
	"os/user"
	"path"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ogier/pflag"

	"github.com/awishformore/logger"

	"github.com/awishformore/m3/adaptor/arbiter"
	"github.com/awishformore/m3/business"
)

const (
	major = 0
	minor = 1
	patch = 1
)

func main() {

	// get current user
	usr, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAILED TO GET CURRENT USER (%v)\n", err)
		os.Exit(1)
	}

	// define configuration parameters
	level := pflag.StringP("level", "l", "INFO", "log level")
	ipc := pflag.StringP("ipc", "i", path.Join(usr.HomeDir, ".ethereum", "testnet", "geth.ipc"), "IPC endpoint for Ethereum node")
	maker := pflag.StringP("maker", "m", "0x5661e7bc2403c7cc08df539e4a8e2972ec256d11", "Maker Market contract address")
	proxy := pflag.StringP("proxy", "p", "0x5661e7bc2403c7cc08df539e4a8e2972ec256d12", "Trade Proxy contract address")
	refresh := pflag.DurationP("refresh", "r", time.Second*14, "interval to check poll for new orders on market")
	threshold := pflag.Uint64P("threshold", "t", 30000, "threshold of profit margin to execute trades")
	pflag.Parse()

	// initialize logger
	lvl := logger.ParseLevel(*level)
	log := logger.New(os.Stderr, lvl)

	log.Infof("M3 DAEMON v%v.%v.%v", major, minor, patch)
	log.Infof("starting m3 daemon...")

	// initialize connection to the geth node
	conn, err := rpc.NewIPCClient(*ipc)
	if err != nil {
		log.Criticalf("could not initialize IPC connection: %v (%v)", ipc, err)
		os.Exit(1)
	}

	// initialize the contract backend for our wrappers
	backend := backends.NewRPCBackend(conn)

	// initialize the wrapper around the market contract
	arb, err := arbiter.NewMaker(
		backend,
		common.HexToAddress(*maker),
		common.HexToAddress(*proxy),
	)
	if err != nil {
		log.Criticalf("could not initialize the blockchain wrapper (%v)", err)
		os.Exit(1)
	}

	// initialize matcher logic
	matcher := business.NewMatcher(
		log, arb, arb, arb,
		business.SetRefresh(*refresh),
		business.SetThreshold(*threshold),
	)

	log.Infof("m3 daemon startup complete")

	// wait for signal to shut down
	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigc

	log.Infof("shutting down m3 daemon...")

	// stop execution & free resources on relevant modules
	matcher.Stop()
	conn.Close()

	log.Infof("m3 daemon shutdown complete")

	os.Exit(0)
}
