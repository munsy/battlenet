package main

import (
	"flag"
	"fmt"
	"os"
)

var tokenFlag = flag.String("t", "", "Set the OAuth2 token for Battle.net profile API access.")
var keyFlag = flag.String("k", "", "Set the Battle.net client key for API access.")

func checkTokenFlag() {
	if *tokenFlag == "" {
		fmt.Println("Token not set.")
		os.Exit(1)
	}
}

func checkKeyFlag() {
	if *keyFlag == "" {
		fmt.Println("Key not set.")
		os.Exit(1)
	}
}

func init() {

}

func main() {
	flag.Parse()

	checkTokenFlag()
	checkKeyFlag()
}
