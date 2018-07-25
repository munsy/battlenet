package sc2

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/models/sc2"
	"github.com/munsy/gobattlenet/pkg/quota"
	"github.com/munsy/gobattlenet/pkg/regions"
)

// Service represents the Starcraft II service.
type Service struct {
	client *http.Client
	region regions.Region
	locale locale.Locale
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
func (s *Service) get(endpoint string, v interface{}) (*Response, error) {
	if nil == v {
		return nil, errors.ErrNoInterfaceSupplied
	}

	response, err := s.client.Get(endpoint + "?locale=" + s.locale.String() + "&apikey=" + s.key)
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

	return &Response{
		Data:     v,
		Endpoint: endpoint,
		Quota:    q,
	}, nil
}

// Character returns the Sc2 character profile.
func (s *Service) Character(name string, profileID int) (*Response, error) {
	var character *sc2.Character

	return s.get(endpointProfile(s.region, profileID, s.region.Itoa(), name), &character)
}

// LadderSeasons returns the Sc2 profile's ladder seasons.
func (s *Service) LadderSeasons(name string, profileID int) (*Response, error) {
	var ladderSeasons *sc2.LadderSeasons

	return s.get(endpointLadderProfile(s.region, profileID, s.region.Itoa(), name), &ladderSeasons)
}

// MatchHistory returns the Sc2 profile's match history.
func (s *Service) MatchHistory(name string, profileID int) (*Response, error) {
	var matchHistory *sc2.MatchHistory

	return s.get(endpointMatchHistory(s.region, profileID, s.region.Itoa(), name), &matchHistory)
}

// Ladder returns Sc2 ladder data.
func (s *Service) Ladder(id int) (*Response, error) {
	var ladder *sc2.Ladder

	return s.get(endpointLadder(s.region, id), &ladder)
}

// Achievements returns Sc2 achievement data.
func (s *Service) Achievements(id int) (*Response, error) {
	var achievements *sc2.AchievementsData

	return s.get(endpointAchievements(s.region), &achievements)
}

// Rewards returns Sc2 reward data.
func (s *Service) Rewards(id int) (*Response, error) {
	var rewards *sc2.RewardsData

	return s.get(endpointRewards(s.region), &rewards)
}
