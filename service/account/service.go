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

// Service represents the account service.
type Service struct {
	client *http.Client
	token  string
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

// New returns a new account service.
func New(token string, c *http.Client) *Service {
	return &Service{
		client: c,
		token:  token,
	}
}

// BattleID returns the Battle.net ID and BattleTag.
func (s *Service) BattleID() (*account.BattleID, error) {
	var bid *account.BattleID

	err := c.get(endpointUser(c.region), &bid)

	if nil != err {
		return nil, err
	}

	return bid, nil
}

// Sc2OauthProfile returns details about a character.
func (s *Service) Sc2OauthProfile() (*sc2.Character, error) {
	var character *sc2.Character

	err := c.get(endpointSc2User(c.region), &character)

	if nil != err {
		return nil, err
	}

	return character, nil
}

// WoWOauthProfile returns details about the WoW account.
func (s *Service) WoWOauthProfile() (*wow.Characters, error) {
	var character *wow.Characters

	err := c.get(endpointWowCharacters(c.region), &character)

	if nil != err {
		return nil, err
	}

	return character, nil
}
