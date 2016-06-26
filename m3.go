package main

import (
	"fmt"
	"os"
	"os/signal"
	"os/user"
	"path"
	"syscall"

	"github.com/ogier/pflag"

	"github.com/awishformore/m3/adaptor"
	"github.com/awishformore/m3/business"
	"github.com/awishformore/m3/infrastructure/logger"
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
	pflag.Parse()

	// initialize logger
	lvl, err := logger.ParseLevel(*level)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAILED TO INITIALIZE LOGGER (%v)\n", err)
		os.Exit(1)
	}
	lgr := logger.New(lvl)

	lgr.Infof("starting m3 daemon...")

	// initialize the blockchain wrapper with the contract addresses
	market, err := adaptor.NewAtomicMarket(*ipc, *maker, *proxy)
	if err != nil {
		lgr.Criticalf("could not initialize the blockchain wrapper (%v)", err)
		os.Exit(1)
	}

	// initialize matcher logic
	matcher, err := business.NewMatcher(market)
	if err != nil {
		lgr.Criticalf("could not initialize market matcher (%v)", err)
		os.Exit(1)
	}

	lgr.Infof("m3 daemon startup complete")

	// wait for signal to shut down
	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigc

	lgr.Infof("shutting down m3 daemon...")

	// stop execution & free resources on relevant modules
	matcher.Stop()
	market.Close()

	lgr.Infof("m3 daemon shutdown complete")

	os.Exit(0)
}
