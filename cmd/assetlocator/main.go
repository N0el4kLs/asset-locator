package main

import (
	"time"

	"github.com/N0el4kLs/asset-locator/runner"

	"github.com/projectdiscovery/gologger"
)

func main() {
	options, err := runner.ParseOptions()
	if err != nil {
		gologger.Fatal().Msgf("Could not parse options: %s\n", err)
	}
	newRunner, err := runner.NewRunner(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	start := time.Now()
	if err = newRunner.Run(); err != nil {
		panic(err)
	}
	gologger.Info().Msgf("Task done,cost: %v\n", time.Since(start))
}
