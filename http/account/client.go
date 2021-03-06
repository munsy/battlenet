package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/battlenet/errors"
	"github.com/munsy/battlenet/http/account/response"
	"github.com/munsy/battlenet/locale"
	"github.com/munsy/battlenet/models/account"
	"github.com/munsy/battlenet/models/sc2"
	"github.com/munsy/battlenet/models/wow"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// Client represents the account client.
type Client struct {
	client    *http.Client
	region    regions.Region
	locale    locale.Locale
	userAgent string
	token     string
}

// Region returns the client's current region as a string.
func (c *Client) Region() string {
	return c.region.String()
}

// Locale returns the client's current locale as a string.
func (c *Client) Locale() string {
	return c.locale.String()
}

// UserAgent returns the client User-Agent header used in API requests.
func (c *Client) UserAgent() string {
	return c.userAgent
}

// New returns a new account client.
func New(c *http.Client, r regions.Region, l locale.Locale, t, v string) (*Client, error) {
	if "" == t {
		return nil, errors.ErrNoTokenSupplied
	}

	if "" == v {
		return nil, errors.ErrNoVersionSupplied
	}

	ac := &Client{
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		region:    regions.US,
		userAgent: "GoBattleNet/" + v,
		token:     t,
	}

	if r.Int() > 5 {
		return nil, errors.ErrUnsupportedArgument
	}
	if nil != c {
		ac.client = c
	}
	if ac.locale != l {
		ac.locale = l
	}
	if ac.region != r {
		ac.region = r
	}

	return ac, nil
}

// BattleID returns the Battle.net ID and BattleTag.
func (c *Client) BattleID() (*response.BattleID, error) {
	var bid *account.BattleID

	ep := endpointUser(c.region)

	q, err := c.get(ep, &bid)

	if nil != err {
		return nil, err
	}

	return &response.BattleID{
		Data:     bid,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Sc2OauthProfile returns details about a character.
func (c *Client) Sc2OauthProfile() (*response.Sc2Character, error) {
	var character *sc2.Character

	ep := endpointSc2User(c.region)

	q, err := c.get(ep, &character)

	if nil != err {
		return nil, err
	}

	return &response.Sc2Character{
		Data:     character,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// WoWOauthProfile returns details about the WoW account.
func (c *Client) WoWOauthProfile() (*response.WoWCharacters, error) {
	var characters *wow.Characters

	ep := endpointWowCharacters(c.region)

	q, err := c.get(ep, &characters)

	if nil != err {
		return nil, err
	}

	return &response.WoWCharacters{
		Data:     characters,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Converts an HTTP response from a given endpoint to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given endpoint and will return an error if it fails to properly unmarshal.
func (c *Client) get(endpoint string, v interface{}) (*quota.Quota, error) {
	if nil == v {
		return nil, errors.ErrNoInterfaceSupplied
	}

	request := endpoint + "?locale=" + c.locale.String() + "&access_token=" + c.token

	response, err := c.client.Get(request)
	if nil != err {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &v)
	if nil != err {
		return nil, err
	}

	q := &quota.Quota{}
	err = q.Set(response)
	if nil != err {
		return nil, err
	}

	return q, nil
}
