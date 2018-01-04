package main

import (
	"asynchronous"
	"fmt"
	"synchronous"

	"github.com/spf13/pflag"
)

var scenario *string

func init() {
	scenario = pflag.StringP("scenario", "s", "1",
		"choose a scenario, 1 for the synchronized scenario,"+
			" 2 for the asynchronized one"+
			" default to 1")
	pflag.Parse()
}

func main() {
	switch *scenario {
	case "1":
		asynchronous.Start()
	case "2":
		synchronous.Start()
	default:
		fmt.Println("Please choose a scenario, 1 for the synchronized one, 2 for the async one")
	}
}
