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

import (
        "flag"
        "fmt"

        "github.com/munsy/gobattlenet/account"
)

var tokenFlag = flag.String("t", "", "Battle.net API token (required).")

func main() {
        flag.Parse()

        client, err := account.New(*tokenFlag)

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

```
```
$ go build
$ ./account -t $YOUR_API_TOKEN_HERE
ID: 12345654321
BattleTag: Munsy#78910
```

## License
MIT
