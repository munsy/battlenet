package battlenet

import (
	"errors"
	"strings"
)

var url = func(s string) string {
	r := strings.ToLower(region)

	if r == "cn" {
		return "https://www.battlenet.com.cn/oauth/" + s
	}
	if r == "us" || r == "eu" || r == "apac" {
		return "https://" + r + ".battle.net/oauth/" + s
	}
	return "https://us.battle.net/oauth/" + s
}

var (
	EndpointOauth2 = func(endpoint string) string {
		if "token" != endpoint && "authorize" != endpoint {
			panic(errors.New("Can't resolve endpoint."))
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

	EndpointAPI = func() string {
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

	EndpointAuthURL  = EndpointOauth2("authorize")
	EndpointTokenURL = EndpointOauth2("token")
)
