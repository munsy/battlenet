package wow

import (
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/regions"
	"github.com/munsy/gobattlenet/settings"
)

// WoWClient allows the user to access the World of Warcraft Battle.net API.
type WoWClient struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	region    regions.Region
	key       string
}

// New creates a new WoWClient. Passing different interface values/types
// can cause different behaviors. See function definition for more details.
func New(args ...interface{}) (c *WoWClient, err error) {
	c = &WoWClient{
		userAgent: "GoBattleNetWow/" + settings.ClientVersion,
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
			return nil, errors.ErrUnsupportedArgument
		}
	}
	return c, nil
}

// Locale gets the client's locale.
func (c *WoWClient) Locale() locale.Locale {
	return c.locale
}

// SetLocale sets the client's locale.
func (c *WoWClient) SetLocale(locale locale.Locale) {
	c.locale = locale
}

// SetKey sets the client's key.
func (c *WoWClient) SetKey(key string) {
	c.key = key
}

// UserAgent returns the client's user agent.
func (c *WoWClient) UserAgent() string {
	return c.userAgent
}
