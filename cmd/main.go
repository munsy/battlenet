package main

import (
	"fmt"
	"os"

	"github.com/munsy/gobattlenet/account"
	"github.com/munsy/gobattlenet/d3"
	"github.com/munsy/gobattlenet/sc2"
	"github.com/munsy/gobattlenet/wow"
)

var config Config

func init() {
	readTOML(*configFlag)
	if *keyFlag == "" {
		*keyFlag = config.Key
	}
	if *tokenFlag == "" {
		*tokenFlag = config.Token
	}
	if *regionFlag == "" {
		*regionFlag = config.Region
	}
	if *localeFlag == "" {
		*localeFlag = config.Locale
	}
}

func main() {
	if len(os.Args) < 2 {
		printDefaultsAndQuit()
	}

	switch os.Args[1] {
	case "config":
		configCommand.Parse(os.Args[2:])
	case "account":
		accountCommand.Parse(os.Args[2:])
	case "d3":
		d3Command.Parse(os.Args[2:])
	case "sc2":
		sc2Command.Parse(os.Args[2:])
	case "wow":
		wowCommand.Parse(os.Args[2:])
	default:
		printDefaultsAndQuit()
	}

	if configCommand.Parsed() {
		parseConfigCommand()
	}
	if accountCommand.Parsed() {
		parseAccountCommand()
	}
	if d3Command.Parsed() {
		parseD3Command()
	}
	if sc2Command.Parsed() {
		parseSC2Command()
	}
	if wowCommand.Parsed() {
		parseWoWCommand()
	}
}

func printDefaultsAndQuit() {
	if len(os.Args) < 2 {
		fmt.Print("[ERROR] No arguments supplied. ")
	} else {
		fmt.Print("[ERROR] Invalid argument. ")
	}
	fmt.Println("Possible arguments are:")
	fmt.Println("\t- config (Configuration settings)")
	fmt.Println("\t- account (Account API access)")
	fmt.Println("\t- d3 (Diablo III API access)")
	fmt.Println("\t- sc2 (Starcraft II API access)")
	fmt.Println("\t- wow (World of Warcraft API access)")
	os.Exit(1)
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

func parseAccountCommand() {
	checkAllAccountConfigs()
	client, err := account.New(config.Settings("token"))

	if nil != err {
		panic(err)
	}

	if *accountBattleIDFlag == true {
		response, err := client.BattleID()

		if nil != err {
			panic(err)
		}

		fmt.Print("ID: %v\n", response.ID)
		fmt.Print("BattleTag: %v\n", response.BattleTag)
	} else if *accountSc2ProfileFlag == true {
		response, err := client.Sc2OauthProfile()

		if nil != err {
			panic(err)
		}

		fmt.Print("%v\n", response)
	} else if *accountWoWProfileFlag == true {
		response, err := client.WoWOauthProfile()

		if nil != err {
			panic(err)
		}

		fmt.Print("%v\n", response)
	} else {
		accountCommand.PrintDefaults()
	}
}

func parseD3Command() {
	checkAllGameConfigs()
	client, err := d3.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Print("%v\n", client)
}

func parseSC2Command() {
	checkAllGameConfigs()
	client, err := sc2.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Print("%v\n", client)
}

func parseWoWCommand() {
	checkAllGameConfigs()
	client, err := wow.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Print("%v\n", client)
}
