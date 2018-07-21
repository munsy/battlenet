package account

import (
	"errors"
	"fmt"

	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/regions"
	"github.com/munsy/gobattlenet/sc2"
	"github.com/munsy/gobattlenet/wow"
)

// AccountClient implements the battlenet.Client interface.
type AccountClient struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	region    battlenet.Region
	token     string
}

// New creates a new AccountClient. Passing different interface types can cause
// different behaviors. See function definiton for more details.
func New(args ...interface{}) (c *AccountClient, err error) {
	c = &AccountClient{
		userAgent: "GoBattleNetAccount/" + battlenet.ClientVersion,
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		region:    battlenet.US,
		token:     "",
	}

	if nil == args {
		return c, nil
	}

	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			c.token = t
			break
		case *http.Client:
			c.client = t
			break
		case locale.Locale:
			c.locale = t
			break
		case battlenet.Region:
			c.region = t
			break
		case battlenet.BNetSettings:
			c.client = t.Client
			c.locale = t.Locale
			c.region = t.Region
			c.token = t.Key
			break
		default:
			return nil, battlenet.ErrorUnsupportedArgument
		}
	}
	return c, nil
}

// Locale gets the client's locale.
func (c *AccountClient) Locale() locale.Locale {
	return c.locale
}

// SetLocale sets the client's locale.
func (c *AccountClient) SetLocale(locale locale.Locale) {
	c.locale = locale
}

// SetKey sets the client's key.
func (c *AccountClient) SetKey(token string) {
	c.token = token
}

// UserAgent returns the client's user agent.
func (c *AccountClient) UserAgent() string {
	return c.userAgent
}
