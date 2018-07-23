package sc2

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/models/sc2"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

// Client allows the user to access the Starcraft II Battle.net API.
type Client struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	region    regions.Region
	key       string
}

// New creates a new Client. Passing different interface values/types
// can cause different behaviors. See function definition for more details.
func New(s *settings.BNetSettings) (c *Client, err error) {
	c = &Client{
		userAgent: "GoBattleNetSC2/" + settings.ClientVersion,
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		region:    regions.US,
		key:       "",
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
		c.key = s.Key
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
func (c *Client) SetKey(key string) {
	c.key = key
}

// UserAgent returns the client's user agent.
func (c *Client) UserAgent() string {
	return c.userAgent
}

// Character returns the Sc2 character profile.
func (c *Client) Character(name string, profileID int) (*sc2.Character, error) {
	var character *sc2.Character

	err := c.get(endpointProfile(c.region, profileID, c.region.Itoa(), name), &character)

	if nil != err {
		return nil, err
	}

	return character, nil
}

// LadderSeasons returns the Sc2 profile's ladder seasons.
func (c *Client) LadderSeasons(name string, profileID int) (*sc2.LadderSeasons, error) {
	var ladderSeasons *sc2.LadderSeasons

	err := c.get(endpointLadderProfile(c.region, profileID, c.region.Itoa(), name), &ladderSeasons)

	if nil != err {
		return nil, err
	}

	return ladderSeasons, nil
}

// MatchHistory returns the Sc2 profile's match history.
func (c *Client) MatchHistory(name string, profileID int) (*sc2.MatchHistory, error) {
	var matchHistory *sc2.MatchHistory

	err := c.get(endpointMatchHistory(c.region, profileID, c.region.Itoa(), name), &matchHistory)

	if nil != err {
		return nil, err
	}

	return matchHistory, nil
}

// Ladder returns Sc2 ladder data.
func (c *Client) Ladder(id int) (*sc2.Ladder, error) {
	var ladder *sc2.Ladder

	err := c.get(endpointLadder(c.region, id), &ladder)

	if nil != err {
		return nil, err
	}

	return ladder, nil
}

// Achievements returns Sc2 achievement data.
func (c *Client) Achievements(id int) (*sc2.AchievementsData, error) {
	var achievements *sc2.AchievementsData

	err := c.get(endpointAchievements(c.region), &achievements)

	if nil != err {
		return nil, err
	}

	return achievements, nil
}

// Rewards returns Sc2 reward data.
func (c *Client) Rewards(id int) (*sc2.RewardsData, error) {
	var rewards *sc2.RewardsData

	err := c.get(endpointRewards(c.region), &rewards)

	if nil != err {
		return nil, err
	}

	return rewards, nil
}

// Convert an HTTP response from a given endpoint to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given endpoint and will return an error if it fails to properly unmarshal.
func (c *Client) get(endpoint string, v interface{}) error {
	if nil == v {
		return errors.ErrNoInterfaceSupplied
	}

	response, err := http.Get(endpoint + "?locale=" + c.locale.String() + "&apikey=" + c.key)
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
