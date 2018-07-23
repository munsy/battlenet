package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/client/wow"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

var keyFlag = flag.String("k", "", "Battle.net API key (required).")

func main() {
	flag.Parse()

	settings := &settings.BNetSettings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
		Key:    *keyFlag,
	}

	client, err := wow.New(settings)

	if nil != err {
		panic(err)
	}

	status, err := client.RealmStatus()

	if nil != err {
		panic(err)
	}

	fmt.Printf("%20s%15s\n", "Realm", "Available?")
	fmt.Println("----------------------------------------")
	for i := 0; i < len(status.Realms); i++ {
		fmt.Printf("%20s%10v\n", status.Realms[i].Name, status.Realms[i].Status)
	}
}
