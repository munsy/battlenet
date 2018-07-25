package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet"
	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/regions"
)

var tokenFlag = flag.String("t", "", "Battle.net API token (required).")

func main() {
	flag.Parse()

	if *tokenFlag == "" {
		fmt.Println("No token provided.")
		return
	}

	settings := &battlenet.Settings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
	}

	client, err := battlenet.AccountClient(settings, *tokenFlag)

	if nil != err {
		panic(err)
	}

	response, err := client.BattleID()

	if nil != err {
		panic(err)
	}

	bid := response.Data

	fmt.Printf("ID: %d\n", bid.ID)
	fmt.Printf("BattleTag: %s\n", bid.BattleTag)
}
