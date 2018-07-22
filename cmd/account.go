package main

import (
	"fmt"
	"os"

	"github.com/munsy/gobattlenet/account"
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

		fmt.Printf("ID: %v\n", response.ID)
		fmt.Printf("BattleTag: %v\n", response.BattleTag)
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
