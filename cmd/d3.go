package main

import (
	"flag"
	"fmt"

	"github.com/munsy/gobattlenet/d3"
)

var (
	d3Command = flag.NewFlagSet("d3", flag.ExitOnError)
)

func parseD3Command() {
	checkAllGameConfigs()
	client, err := d3.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Print("%v\n", client)
}
