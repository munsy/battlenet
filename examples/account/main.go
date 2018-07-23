package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/client/account"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

var tokenFlag = flag.String("t", "", "Battle.net API token (required).")

func main() {
	flag.Parse()

	settings := &settings.BNetSettings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
		Key:    *tokenFlag,
	}

	client, err := account.New(settings)

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
