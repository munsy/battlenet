package main

import (
	"flag"
	"fmt"
	"os"
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

	if accountCommand.Parsed() {
		// Required Flags
		if *accountBattleIDFlag == true {
			// call account api and get battleid
		} else if *accountSc2ProfileFlag == true {
			// call account api and get sc2 profile
		} else if *accountWoWProfileFlag == true {
			// call account api and get wow profile
		} else {
			accountCommand.PrintDefaults()
		}
	}

}
