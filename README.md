# gobattlenet

Bindings for Battle.net's API written in Golang.

## Install
```
$ go get github.com/munsy/gobattlenet/account
$ go get github.com/munsy/gobattlenet/d3
$ go get github.com/munsy/gobattlenet/sc2
$ go get github.com/munsy/gobattlenet/wow
```

## Usage
```
package main

import(
        "fmt"
        "os"

        "github.com/munsy/gobattlenet/account"
)

func main() {
		key := os.Args[1]
        client := account.New()

        character, err := client.GetCharacterProfile("us", "thrall", "munsy", "")

        if nil != err {
                panic(err)
        }

        fmt.Printf("Name: %v\nClass: %v\nRace: %v\n", character.Name, character.Class, character.Race)
}
```
```
$ go build
$ ./raideriotest
Name: Munsy
Class: Druid
Race: Troll
```

## Disclaimer
I am in no way associated with the fine folks over at https://www.raider.io other than being a fan of their work and wanting to use it with my own stuff.
