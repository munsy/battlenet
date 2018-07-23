package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/service"
	"github.com/munsy/gobattlenet/service/account"
	"github.com/munsy/gobattlenet/service/d3"
	"github.com/munsy/gobattlenet/service/sc2"
	"github.com/munsy/gobattlenet/service/wow"
	"github.com/munsy/gobattlenet/settings"
)

const clientVersion = "alpha"

// Client allows the user to access the Battle.net API.
type Client struct {
	Account   *account.Service
	DIII      *d3.Service
	ScII      *sc2.Service
	WoW       *wow.Service
	client    *http.Client
	userAgent string
	locale    locale.Locale
	region    regions.Region
}

// New creates a new Client.
func New(s *settings.BNetSettings) (c *Client, err error) {
	c = &Client{
		Account:   service.Account(),
		DIII:      service.D3(),
		ScII:      service.Sc2(),
		WoW:       service.WoW(),
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
	if s.Key != "" {
		c.token = s.Key
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

// SetKey sets the client's key.
func (c *Client) SetKey(token string) {
	c.token = token
}

// UserAgent returns the client's user agent.
func (c *Client) UserAgent() string {
	return c.userAgent
}

// Converts an HTTP response from a given endpoint to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given endpoint and will return an error if it fails to properly unmarshal.
func (s *Service) get(endpoint string, v interface{}) error {
	if nil == v {
		return errors.ErrNoInterfaceSupplied
	}

	response, err := c.client.Get(endpoint + "?access_token=" + c.token)
	if nil != err {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err
	}

	err = json.Unmarshal([]byte(body), &v)
	if nil != err {
		return err
	}

	return nil
}
