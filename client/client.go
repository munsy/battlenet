package client

import (
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/service/account"
	"github.com/munsy/gobattlenet/service/d3"
	"github.com/munsy/gobattlenet/service/sc2"
	"github.com/munsy/gobattlenet/service/wow"
	"github.com/munsy/gobattlenet/settings"
)

const clientVersion = "alpha"

// Client allows the user to access the Battle.net API.
type Client struct {
	client    *http.Client
	userAgent string
	locale    locale.Locale
	region    regions.Region
}

// New creates a new Client.
func New(s *settings.BNetSettings) (c *Client, err error) {
	c = &Client{
		client:    &http.Client{Timeout: (10 * time.Second)},
		userAgent: "GoBattleNet/" + clientVersion,
		locale:    locale.AmericanEnglish,
		region:    regions.US,
	}

	if nil == s {
		return c, nil
	}
	if s.Region.Int() > 5 {
		return nil, errors.ErrUnsupportedArgument
	}
	if nil != s.Client {
		c.client = s.Client
	}
	if c.locale != s.Locale {
		c.locale = s.Locale
	}
	if c.region != s.Region {
		c.region = s.Region
	}

	return c, nil
}

// Locale gets the client's locale.
func (c *Client) Locale() locale.Locale {
	return c.locale
}

// SetLocale sets the client's locale.
func (c *Client) SetLocale(locale locale.Locale) {
	c.locale = locale
}

// UserAgent returns the client's user agent.
func (c *Client) UserAgent() string {
	return c.userAgent
}

// Account returns a new Account service.
func (c *Client) Account(token string) *account.Service {
	return account.New(token, c.client)
}

// D3 returns a new D3 service.
func (c *Client) D3(key string) *d3.Service {
	return d3.New(key, c.client)
}

// Sc2 returns a new Sc2 service.
func (c *Client) Sc2(key string) *sc2.Service {
	return sc2.New(key, c.client)
}

// WoW returns a new WoW service.
func (c *Client) WoW(key string) *wow.Service {
	return wow.New(key, c.client)
}
