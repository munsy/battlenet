package main

import (
	"flag"
	"fmt"

	"github.com/munsy/gobattlenet/d3"
)

var keyFlag = flag.String("k", "", "Battle.net API key (required).")

func main() {
	flag.Parse()

	client, err := d3.New(*keyFlag)

	if nil != err {
		panic(err)
	}

	acts, err := client.ActIndex()

	if nil != err {
		panic(err)
	}

	fmt.Printf("%s\t\t%s\n", "Number", "Name")
	fmt.Println("-----------------------")
	for i := 0; i < len(acts.Acts); i++ {
		fmt.Printf("%d\t\t%v\n", acts.Acts[i].Number, acts.Acts[i].Name)
	}
}
