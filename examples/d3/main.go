package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/client/d3"
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
	}

	client, err := http.New(settings)

	if nil != err {
		panic(err)
	}

	acts, err := client.D3(*keyFlag).ActIndex()

	if nil != err {
		panic(err)
	}

	fmt.Printf("%s\t\t%s\n", "Number", "Name")
	fmt.Println("-----------------------")
	for i := 0; i < len(acts.Acts); i++ {
		fmt.Printf("%d\t\t%v\n", acts.Acts[i].Number, acts.Acts[i].Name)
	}
}
