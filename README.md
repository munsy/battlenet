GoBattleNet
===========
[![Build Status](https://travis-ci.org/munsy/gobattlenet.svg?branch=master)](https://travis-ci.org/Munsy/gobattlenet) [![Go Report Card](https://goreportcard.com/badge/github.com/munsy/gobattlenet)](https://goreportcard.com/report/github.com/munsy/gobattlenet) [![GoDoc](https://godoc.org/github.com/munsy/gobattlenet?status.svg)](https://godoc.org/github.com/munsy/gobattlenet)

Bindings for Battle.net's API written in Golang.

## Versions
* This library is currently in alpha. Everything is subject to change. Use at your own discretion.
* All data returned by clients should be up-to-date, including 
World of Warcraft, which just had its API updated for Battle for Azeroth.

## Install
```
$ go get github.com/munsy/gobattlenet
```
## Example
```
package main

import (
        "flag"
        "fmt"

        "github.com/munsy/gobattlenet"
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

        client, err := battlenet.NewAccountClient(settings)

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
$ go build -o account.exe
$ ./account.exe -t $YOUR_API_TOKEN_HERE
ID: 12345654321
BattleTag: Munsy#78910
```

Additional examples for Diablo III, Starcraft II, and World of Warcraft can be 
found in the [examples](https://github.com/munsy/gobattlenet/blob/master/examples) 
package.

Bug Reporting/Fixing
====================
Please submit all bug reports as issues.
This section will be updated in the future with a more comprehensive guide to 
submitting helpful bug reports and/or bug fixes.

Licensing
=========
GoBattleNet is licensed under the MIT License. See
[LICENSE](https://github.com/munsy/gobattlenet/blob/master/LICENSE) for the full
license text.