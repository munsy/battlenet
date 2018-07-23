package main

import (
	"flag"
	"fmt"

	"github.com/munsy/gobattlenet/client/account"
)

var tokenFlag = flag.String("t", "", "Battle.net API token (required).")

func main() {
	flag.Parse()

	client, err := account.New(*tokenFlag)

	if nil != err {
		panic(err)
	}

	bid, err := client.BattleID()

	if nil != err {
		panic(err)
	}

	fmt.Printf("ID: %d\n", bid.ID)
	fmt.Printf("BattleTag: %s\n", bid.BattleTag)
}
