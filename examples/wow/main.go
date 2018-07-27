package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/battlenet"
)

var keyFlag = flag.String("k", "", "Battle.net API key (required).")

func main() {
	flag.Parse()

	settings := &battlenet.Settings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: battlenet.Locale.AmericanEnglish,
		Region: battlenet.Regions.US,
	}

	client, err := battlenet.WoWClient(settings, *keyFlag)

	if nil != err {
		panic(err)
	}

	response, err := client.RealmStatus()

	if nil != err {
		panic(err)
	}

	status := response.Data

	fmt.Printf("%20s%15s\n", "Realm", "Available?")
	fmt.Println("----------------------------------------")
	for i := 0; i < len(status.Realms); i++ {
		fmt.Printf("%20s%10v\n", status.Realms[i].Name, status.Realms[i].Status)
	}
}
