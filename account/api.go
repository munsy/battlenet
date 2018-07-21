package account

import (
	"github.com/munsy/gobattlenet/sc2"
	"github.com/munsy/gobattlenet/wow"
)

/*
	Account related API methods should go in here.
*/

// BattleID returns the Battle.net ID and BattleTag.
func (c *AccountClient) BattleID(region string) (*BattleID, error) {
	var bid *BattleID

	err := c.get(endpointUser(c.region), bid)

	if nil != err {
		return nil, err
	}

	return bid, nil
}

// Sc2OauthProfile returns details about a character.
func (c *AccountClient) Sc2OauthProfile(region string) (*sc2.Character, error) {
	var character *sc2.Character

	err := c.get(endpointSc2User(c.region), character)

	if nil != err {
		return nil, err
	}

	return character, nil
}

// WoWOauthProfile returns details about the WoW account.
func (c *AccountClient) WoWOauthProfile() (*wow.Characters, error) {
	var character *wow.Characters

	err := c.get(endpointWowCharacters(c.region), character)

	if nil != err {
		return nil, err
	}

	return character, nil
}
