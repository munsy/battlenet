package battlenet

type Region int

const (
	US Region = iota
	EU
	KR
	TW
	SEA
	CN
)

func (r Region) Int() int {
	return int(r)
}

func (r Region) String() string {
	var region string
	switch r {
	case US:
		region = "us"
	case EU:
		region = "eu"
	case KR:
		region = "kr"
	case TW:
		region = "tw"
	case SEA:
		region = "sea"
	case CN:
		region = "cn"
	default:
		panic(ErrorUnresolvedEndpoint)
	}
	return region
}

// AuthURL returns the URL for OAuth authorization.
func (r Region) AuthURL() string {
	return r.oauthHelper("authorize")
}

// TokenURL returns the URL to retrieve an OAuth token.
func (r Region) TokenURL() string {
	return r.oauthHelper("token")
}

// API returns the URL relative to the region for REST calls.
func (r Region) API() string {
	var url string
	switch r {
	case US:
		url = "https://" + r.String() + ".api.battle.net/"
		break
	case EU:
		url = "https://" + r.String() + ".api.battle.net/"
		break
	case KR:
		url = "https://" + r.String() + ".api.battle.net/"
		break
	case TW:
		url = "https://" + r.String() + ".api.battle.net/"
		break
	case SEA:
		url = "https://" + r.String() + ".api.battle.net/"
		break
	case CN:
		url = "https://api.battlenet.com.cn/"
		break
	default:
		panic(ErrorUnresolvedEndpoint)
	}
	return url
}

func (r Region) oauthHelper(endpoint string) string {
	var url string
	switch r {
	case US:
		url = "https://" + r.String() + ".battle.net/oauth/" + endpoint
		break
	case EU:
		url = "https://" + r.String() + ".battle.net/oauth/" + endpoint
		break
	case KR:
		url = "https://" + r.String() + ".battle.net/oauth/" + endpoint
		break
	case TW:
		url = "https://" + r.String() + ".battle.net/oauth/" + endpoint
		break
	case SEA:
		url = "https://" + r.String() + ".battle.net/oauth/" + endpoint
		break
	case CN:
		url = "https://www.battlenet.com.cn/oauth/" + endpoint
		break
	default:
		panic(ErrorUnresolvedEndpoint)
	}
	return url
}
