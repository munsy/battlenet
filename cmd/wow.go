package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/munsy/gobattlenet/wow"
)

var (
	wowCommand = flag.NewFlagSet("wow", flag.ExitOnError)
	wowFlag    = wowCommand.Bool("bid", false, "Description.")
)

func parseWoWCommand() {
	checkAllGameConfigs()
	client, err := wow.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Printf("%v\n", client)
}

func printWoWCommandsAndQuit() {
	if len(os.Args) < 3 {
		fmt.Println("[ERROR] No arguments supplied.")
	} else {
		fmt.Print("[ERROR] Invalid argument. ")
	}
	fmt.Println("Possible arguments are:")
	wowCommand.PrintDefaults()
	os.Exit(1)
}
