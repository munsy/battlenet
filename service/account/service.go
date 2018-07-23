package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/models/account"
	"github.com/munsy/gobattlenet/pkg/models/sc2"
	"github.com/munsy/gobattlenet/pkg/models/wow"
	"github.com/munsy/gobattlenet/pkg/quota"
	"github.com/munsy/gobattlenet/pkg/regions"
)

// Service represents the account service.
type Service struct {
	client *http.Client
	region regions.Region
	locale locale.Locale
	token  string
}

// New returns a new account service.
func New(token string, c *http.Client) *Service {
	return &Service{
		client: c,
		token:  token,
	}
}

// BattleID returns the Battle.net ID and BattleTag.
func (s *Service) BattleID() (*Response, error) {
	var bid *account.BattleID
	return s.get(endpointUser(s.region), &bid)
}

// Sc2OauthProfile returns details about a character.
func (s *Service) Sc2OauthProfile() (*Response, error) {
	var character *sc2.Character
	return s.get(endpointSc2User(s.region), &character)
}

// WoWOauthProfile returns details about the WoW account.
func (s *Service) WoWOauthProfile() (*Response, error) {
	var character *wow.Characters
	return s.get(endpointWowCharacters(s.region), &character)
}

// Converts an HTTP response from a given endpoint to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given endpoint and will return an error if it fails to properly unmarshal.
func (s *Service) get(endpoint string, v interface{}) (*Response, error) {
	if nil == v {
		return nil, errors.ErrNoInterfaceSupplied
	}

	response, err := s.client.Get(endpoint + "?access_token=" + s.token)
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
