package d3

import (
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/regions"
	"github.com/munsy/gobattlenet/settings"
)

// D3Client allows the user to access the Diablo III Battle.net API.
type D3Client struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	region    regions.Region
	key       string
}

// New creates a new D3Client. Passing different interface values/types
// can cause different behaviors. See function definition for more details.
func New(args ...interface{}) (c *D3Client, err error) {
	c = &D3Client{
		userAgent: "GoBattleNetD3/" + settings.ClientVersion,
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		region:    regions.US,
		key:       "",
	}

	if nil == args {
		return c, nil
	}

	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			c.key = t
			break
		case *http.Client:
			c.client = t
			break
		case locale.Locale:
			c.locale = t
			break
		case regions.Region:
			c.region = t
			break
		case settings.BNetSettings:
			c.client = t.Client
			c.locale = t.Locale
			c.region = t.Region
			c.key = t.Key
			break
		default:
			return nil, errors.UnsupportedArgument
		}
	}
	return c, nil
}

// Locale gets the client's locale.
func (c *D3Client) Locale() locale.Locale {
	return c.locale
}

// SetLocale sets the client's locale.
func (c *D3Client) SetLocale(locale locale.Locale) {
	c.locale = locale
}

// SetKey sets the client's key.
func (c *D3Client) SetKey(key string) {
	c.key = key
}

// UserAgent returns the client's user agent.
func (c *D3Client) UserAgent() string {
	return c.userAgent
}
