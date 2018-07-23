package main

import (
	"flag"
	"fmt"

	"github.com/munsy/gobattlenet/client/wow"
)

var keyFlag = flag.String("k", "", "Battle.net API key (required).")

func main() {
	flag.Parse()

	client, err := wow.New(*keyFlag)

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
