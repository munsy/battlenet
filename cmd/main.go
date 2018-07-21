package main

import (
	"flag"
	"log"
)

var config Config

func init() {
	readTOML(*configFlag)
	if *keyFlag == "" {
		*keyFlag = config.Key
	}
	if *tokenFlag == "" {
		*tokenFlag = config.Token
	}
}

func main() {
	flag.Parse()

	checkTokenFlag()
	checkKeyFlag()

	if *writeFlag == true {
		log.Fatal(writeTOML(*keyFlag, *tokenFlag))
	}

	return // more soon
}
