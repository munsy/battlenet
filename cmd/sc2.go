package main

import (
	"flag"
	"fmt"

	"github.com/munsy/gobattlenet/sc2"
)

var (
	sc2Command = flag.NewFlagSet("sc2", flag.ExitOnError)
)

func parseSC2Command() {
	checkAllGameConfigs()
	client, err := sc2.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Print("%v\n", client)
}
