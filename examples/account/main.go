package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/models/account"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

var tokenFlag = flag.String("t", "", "Battle.net API token (required).")

func main() {
	flag.Parse()

	if *tokenFlag == "" {
		fmt.Println("No token provided.")
		return
	}

	settings := &settings.BNetSettings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
	}

	client, err := battlenet.New(settings)

	if nil != err {
		panic(err)
	}

	response, err := client.Account(*tokenFlag).BattleID()

	if nil != err {
		panic(err)
	}

	data, err := response.Data()

	if nil != err {
		panic(err)
	}

	bid := data.(*account.BattleID)

	fmt.Printf("ID: %d\n", bid.ID)
	fmt.Printf("BattleTag: %s\n", bid.BattleTag)
}
