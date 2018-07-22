package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	configCommand = flag.NewFlagSet("config", flag.ExitOnError)
	configFlag    = configCommand.String("f", "config.toml", "Configuration file to load key/token values.")
	keyFlag       = configCommand.String("k", "", "Set the Battle.net client key for API access.")
	tokenFlag     = configCommand.String("t", "", "Set the OAuth2 token for Battle.net profile API access.")
	writeFlag     = configCommand.Bool("o", false, "Write current key/token settings in memory to config file.")
	regionFlag    = configCommand.String("r", "", "Set the Battle.net region.") // ADD
	localeFlag    = configCommand.String("l", "", "Set the Battle.net locale.") // ADD
)

func parseConfigCommand() {
	if len(os.Args) < 3 {
		printConfigCommandsAndQuit()
	}

	if *regionFlag != "" {
		config.Region = *regionFlag
	}
	if *localeFlag != "" {
		config.Locale = *localeFlag
	}

	if *writeFlag == true {
		checkAllConfigs()

		err := writeTOML(*keyFlag, *tokenFlag, *regionFlag, *localeFlag)
		if nil != err {
			fmt.Println("Write failed: " + err.Error())
		} else {
			fmt.Println("Write successful.")
		}
	} else {
		printConfigCommandsAndQuit()
	}
}

func printConfigCommandsAndQuit() {
	if len(os.Args) < 3 {
		fmt.Println("[ERROR] No arguments supplied.")
	} else {
		fmt.Print("[ERROR] Invalid argument. ")
	}
	fmt.Println("Possible arguments are:")
	configCommand.PrintDefaults()
	os.Exit(1)

}
