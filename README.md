BattleNet
=========
[![Build Status](https://travis-ci.org/munsy/battlenet.svg?branch=master)](https://travis-ci.org/munsy/battlenet) [![Go Report Card](https://goreportcard.com/badge/github.com/munsy/battlenet)](https://goreportcard.com/report/github.com/munsy/battlenet) [![GoDoc](https://godoc.org/github.com/munsy/battlenet?status.svg)](https://godoc.org/github.com/munsy/battlenet)

A Golang library for retrieving data from Blizzard's Battle.net API.

Requires Go 1.12 or higher.

## Install
```
$ go get github.com/munsy/battlenet
```
## Example
```
package main

import (
        "flag"
        "fmt"
        "net/http"
        "time"

        "github.com/munsy/battlenet"
)

var tokenFlag = flag.String("t", "", "Battle.net API token (required).")

func main() {
        flag.Parse()

        if *tokenFlag == "" {
                fmt.Println("No token provided.")
                return
        }

        // Create settings for the client. This is not required, but
        // is necessary for non-default settings.
        settings := &battlenet.Settings{
                Client: &http.Client{Timeout: (10 * time.Second)},
                Locale: battlenet.Locale.AmericanEnglish,
                Region: battlenet.Regions.US,
        }

        // Create a new client for accessing the Battle.net Account API.
        // There are also clients for Diablo III, Starcraft II, and WoW.
        client, err := battlenet.AccountClient(settings, *tokenFlag)

        if nil != err {
                panic(err)
        }

        // Make a request. Each method corresponds to a Battle.net endpoint.
        response, err := client.BattleID()

        if nil != err {
                panic(err)
        }

        // Get the underlying data. You can also see the endpoint that was called,
        // as well as your quota usage.
        bid := response.Data

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
found in the [examples](https://github.com/munsy/battlenet/blob/master/examples) 
package.

Bug Reporting/Fixing
====================
Please submit all bug reports as issues.
This section will be updated in the future with a more comprehensive guide to 
submitting helpful bug reports and/or bug fixes.

Contributing
============
* Testing is mostly what is needed right now
* Submit contributions as pull requests

Licensing
=========
GoBattleNet is licensed under the MIT License. See
[LICENSE](https://github.com/munsy/battlenet/blob/master/LICENSE) for the full
license text.
