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

// Service represents the Starcraft II service.
type Service struct {
	client *http.Client
	key    string
}

// New returns a new Starcraft II service.
func New(key string, c *http.Client) *Service {
	return &Service{
		client: c,
		key:    key,
	}
}

// Converts an HTTP response from a given endpoint to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given endpoint and will return an error if it fails to properly unmarshal.
func (s *Service) get(endpoint string, v interface{}) error {
	if nil == v {
		return errors.ErrNoInterfaceSupplied
	}

	response, err := c.client.Get(endpoint + "?locale=" + s.locale + "&apikey=" + s.key)
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
