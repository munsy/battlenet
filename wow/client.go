package wow

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet"
	"github.com/munsy/gobattlenet/locale"
)

// WowClient implements the battlenet.Client interface.
type WoWClient struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	region    battlenet.Region
	key       string
}

// New creates a new WoWClient. Passing different interface values/types
// can cause different behaviors. See function definiton for more details.
func New(args ...interface{}) (c *WoWClient, err error) {
	c = &WoWClient{
		userAgent: "GoBattleNetWow/" + battlenet.ClientVersion,
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		region:    battlenet.US,
		key:       "",
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
			c.key = t.Key
			break
		default:
			return nil, battlenet.ErrorUnsupportedArgument
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