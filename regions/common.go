package regions

import (
	"errors"
	"strings"

	"github.com/munsy/gobattlenet"
)

// Region can be one of: "us", "eu", "kr", "tw", "sea", or "cn"
var region string

func SetRegion(r string) {
	region = r
}

var (
	endpointOauth2 = func(endpoint string) string {
		if "token" != endpoint && "authorize" != endpoint {
			panic(battlenet.ErrorUnresolvedEndpoint)
		}

		r := strings.ToLower(region)

		switch r {
		case "us":
		case "eu":
		case "kr":
		case "tw":
		case "sea":
			return "https://" + r + ".battle.net/oauth/" + endpoint
		case "cn":
			return "https://www.battlenet.com.cn/oauth/" + endpoint
		default:
			return "https://us.battle.net/oauth/" + endpoint
		}
		return "https://us.battle.net/oauth/" + endpoint
	}

	API = func() string {
		r := strings.ToLower(region)
		switch r {
		case "us":
		case "eu":
		case "kr":
		case "tw":
		case "sea":
			return "https://" + r + ".api.battle.net/"
		case "cn":
			return "https://api.battlenet.com.cn/"
		default:
			return "https://us.api.battle.net/"
		}
		return "https://us.api.battle.net/"
	}

	endpointData = API() + "data/"

	AuthURL  = endpointOauth2("authorize")
	TokenURL = endpointOauth2("token")
)
