package main

import (
	"flag"
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
		flag.PrintDefaults()
		os.Exit(1)
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
		flag.PrintDefaults()
		os.Exit(1)
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

func parseConfigCommand() {
	if *writeFlag == true {
		checkTokenFlag()
		checkKeyFlag()
		checkRegionFlag()
		checkLocaleFlag()

		err := writeTOML(*keyFlag, *tokenFlag, *regionFlag, *localeFlag)
		if nil != err {
			fmt.Println("Write failed: " + err.Error())
		} else {
			fmt.Println("Write successful.")
		}
	}
}

func parseAccountCommand() {
	client, err := account.New(config.Settings("token"))

	if nil != err {
		panic(err)
	}

	if *accountBattleIDFlag == true {
		response, err := client.BattleID()

		if nil != err {
			panic(err)
		}

		fmt.Print("%v\n", response)
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
	client, err := d3.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Print("%v\n", client)
}

func parseSC2Command() {
	client, err := sc2.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Print("%v\n", client)
}

func parseWoWCommand() {
	client, err := wow.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	fmt.Print("%v\n", client)
}
