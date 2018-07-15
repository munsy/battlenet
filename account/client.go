package account

import (
	"errors"
	"fmt"

	"github.com/munsy/gobattlenet"
	"github.com/munsy/gobattlenet/regions"
	"github.com/munsy/gobattlenet/sc2"
	"github.com/munsy/gobattlenet/wow"
)

// AccountClient implements the battlenet.Client interface.
type AccountClient struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	token     string
}

// New creates a new AccountClient. Passing different interface types can cause
// different behaviors. See function definiton for more details.
func New(args ...interface{}) (c *AccountClient, err error) {
	c = &AccountClient{
		userAgent: "GoBattleNetAccount/" + battlenet.ClientVersion,
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
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
		case battlenet.BNetSettings:
			c.client = t.Client
			c.locale = t.Locale
			c.token = t.Token
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

func (c *AccountClient) BattleID(region string) (*BattleID, error) {
	var bid *BattleID

	regions.SetRegion(region)

	err := c.get(regions.EndpointUser, bid)

	if nil != err {
		return nil, err
	}

	return bid, nil
}

func (c *AccountClient) Sc2OauthProfile(region string) (*sc2.Character, error) {
	var character *sc2.Character

	regions.SetRegion(region)

	err := c.get(endpointSc2User, character)

	if nil != err {
		return nil, err
	}

	return character, nil
}

func (c *AccountClient) WoWOauthProfile() (*wow.Characters, error) {
	var character *wow.Character

	regions.SetRegion(region)

	err := c.get(endpointWowCharacters, character)

	if nil != err {
		return nil, err
	}

	return character, nil
}
