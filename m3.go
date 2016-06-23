package main

import (
	"fmt"
	"os"

	"github.com/ogier/pflag"

	"github.com/awishformore/m3/adaptor/logger"
)

const (
	version string = "0.1.0"
)

func main() {

	// define configuration parameters
	level := pflag.StringP("level", "l", "INFO", "log level")
	pflag.Parse()

	// initialize logger
	lvl, err := logger.ParseLevel(*level)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAILED TO INITIALIZE LOGGER (%v)\n", err)
		os.Exit(1)
	}
	lgr := logger.New(lvl)

	lgr.Infof("starting m3 daemon version %v", version)

	lgr.Infof("shutting down m3 daemon")

	os.Exit(0)
}
