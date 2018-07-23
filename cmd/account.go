package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/munsy/gobattlenet/client/account"
)

var (
	accountCommand        = flag.NewFlagSet("account", flag.ExitOnError)
	accountBattleIDFlag   = accountCommand.Bool("bid", false, "Display Battle.net ID and BattleTag.")
	accountSc2ProfileFlag = accountCommand.Bool("sc2", false, "Display details about a Starcraft II character.")
	accountWoWProfileFlag = accountCommand.Bool("wow", false, "Display details about a World of Warcraft profile.")
)

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
		if nil == response {
			fmt.Println("Nil response received.")
			os.Exit(1)
		}

		printResult(response)
	} else if *accountSc2ProfileFlag == true {
		response, err := client.Sc2OauthProfile()

		if nil != err {
			panic(err)
		}

		printResult(response)
	} else if *accountWoWProfileFlag == true {
		response, err := client.WoWOauthProfile()

		if nil != err {
			panic(err)
		}

		printResult(response)
	} else {
		accountCommand.PrintDefaults()
	}
}
