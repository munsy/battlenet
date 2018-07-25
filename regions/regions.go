package regions

import (
	"github.com/munsy/gobattlenet/errors"
)

// Region can be one of: US, EU, KR, TW, SEA, or CN
type Region int

// String returns the region as a lowercase string.
// If the region is invalid, the function causes a panic.
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
		panic(errors.ErrUnresolvedEndpoint)
	}
	return region
}

// Int returns the region as an integer value.
func (r Region) Int() int {
	return int(r)
}

// Itoa returns the region's integer value as a string representation.
func (r Region) Itoa() string {
	return string(r.Int())
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
		url = urlHead + r.String() + urlTail
		break
	case EU:
		url = urlHead + r.String() + urlTail
		break
	case KR:
		url = urlHead + r.String() + urlTail
		break
	case TW:
		url = urlHead + r.String() + urlTail
		break
	case SEA:
		url = urlHead + r.String() + urlTail
		break
	case CN:
		url = urlCN
		break
	default:
		panic(errors.ErrUnresolvedEndpoint)
	}
	return url
}

// Helper for building OAuth2 URLs.
func (r Region) oauthHelper(endpoint string) string {
	var url string
	switch r {
	case US:
		url = urlHead + r.String() + oauthTail + endpoint
		break
	case EU:
		url = urlHead + r.String() + oauthTail + endpoint
		break
	case KR:
		url = urlHead + r.String() + oauthTail + endpoint
		break
	case TW:
		url = urlHead + r.String() + oauthTail + endpoint
		break
	case SEA:
		url = urlHead + r.String() + oauthTail + endpoint
		break
	case CN:
		url = oauthCN + endpoint
		break
	default:
		panic(errors.ErrUnresolvedEndpoint)
	}
	return url
}
