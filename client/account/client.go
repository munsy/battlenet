package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/models/account"
	"github.com/munsy/gobattlenet/pkg/models/sc2"
	"github.com/munsy/gobattlenet/pkg/models/wow"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

// Client allows the user to access the Battle.net Account API.
type Client struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	region    regions.Region
	token     string
}

// New creates a new Client. Passing different interface types can cause
// different behaviors. See function definition for more details.
func New(args ...interface{}) (c *Client, err error) {
	c = &Client{
		userAgent: "GoBattleNetAccount/" + settings.ClientVersion,
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		region:    regions.US,
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
		case regions.Region:
			c.region = t
			break
		case settings.BNetSettings:
			c.client = t.Client
			c.locale = t.Locale
			c.region = t.Region
			c.token = t.Key
			break
		default:
			return nil, errors.ErrUnsupportedArgument
		}
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

// BattleID returns the Battle.net ID and BattleTag.
func (c *Client) BattleID() (*account.BattleID, error) {
	var bid *account.BattleID

	err := c.get(endpointUser(c.region), &bid)

	if nil != err {
		return nil, err
	}

	return bid, nil
}

// Sc2OauthProfile returns details about a character.
func (c *Client) Sc2OauthProfile() (*sc2.Character, error) {
	var character *sc2.Character

	err := c.get(endpointSc2User(c.region), &character)

	if nil != err {
		return nil, err
	}

	return character, nil
}

// WoWOauthProfile returns details about the WoW account.
func (c *Client) WoWOauthProfile() (*wow.Characters, error) {
	var character *wow.Characters

	err := c.get(endpointWowCharacters(c.region), &character)

	if nil != err {
		return nil, err
	}

	return character, nil
}

// Convert an HTTP response from a given endpoint to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given endpoint and will return an error if it fails to properly unmarshal.
func (c *Client) get(endpoint string, v interface{}) error {
	if nil == v {
		return errors.ErrNoInterfaceSupplied
	}

	response, err := http.Get(endpoint + "?access_token=" + c.token)
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
