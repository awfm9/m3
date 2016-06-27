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

	"github.com/ogier/pflag"

	"github.com/awishformore/logger"

	"github.com/awishformore/m3/adaptor"
	"github.com/awishformore/m3/business"
)

const (
	version string = "0.1.0"
)

func main() {

	// print version info
	fmt.Fprintf(os.Stderr, "M3 DAEMON V%v\n", version)

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
	interval := pflag.DurationP("interval", "t", time.Second*14, "interval to check poll for new orders on market")
	pflag.Parse()

	// initialize logger
	lvl := logger.ParseLevel(*level)
	log := logger.New(lvl)

	log.Infof("starting m3 daemon...")

	// initialize the blockchain wrapper with the contract addresses
	market, err := adaptor.NewAtomicMarket(*ipc, *maker, *proxy)
	if err != nil {
		log.Criticalf("could not initialize the blockchain wrapper (%v)", err)
		os.Exit(1)
	}

	// initialize matcher logic
	matcher := business.NewMatcher(log, market, *interval)

	log.Infof("m3 daemon startup complete")

	// wait for signal to shut down
	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigc

	log.Infof("shutting down m3 daemon...")

	// stop execution & free resources on relevant modules
	matcher.Stop()
	market.Close()

	log.Infof("m3 daemon shutdown complete")

	os.Exit(0)
}
