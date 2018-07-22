package main

import (
	"flag"
	"fmt"

	"github.com/munsy/gobattlenet/wow"
)

var (
	wowCommand = flag.NewFlagSet("wow", flag.ExitOnError)
)

func parseWoWCommand() {
	checkAllGameConfigs()
	client, err := wow.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Print("%v\n", client)
}
