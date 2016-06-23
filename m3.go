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
	level := pflag.StringP("level", "l", "INFO", "leg level")
	pflag.Parse()

	// initialize logger
	logLevel, err := logger.ParseLevel(*level)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAILED TO INITIALIZE LOGGER: %v", err)
		os.Exit(1)
	}
	logger := logger.New(logLevel)

	logger.Infof("starting m3 daemon version %v", version)

	logger.Infof("shutting down m3 daemon")

	os.Exit(0)
}
