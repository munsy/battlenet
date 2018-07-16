package wow

/*
	World of Warcraft related API methods should go in here.
*/

// ACHIEVEMENT API
func (c *WoWClient) Achievement(id int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// AUCTION API
func (c *WoWClient) AuctionDataStatus(realm string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// BOSS API
func (c *WoWClient) BossMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) Boss(bossID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CHALLENGE MODE API
func (c *WoWClient) RealmLeaderboard(realm string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) RegionLeaderboard() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CHARACTER PROFILE API
func (c *WoWClient) CharacterProfile(characterName string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// GUILD PROFILE API
func (c *WoWClient) GuildProfile(guildName string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ITEM API
func (c *WoWClient) Item(itemID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) ItemSet(steID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// MOUNT API
func (c *WoWClient) MountMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PET API
func (c *WoWClient) PetMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) PetAbilities(abilityID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) PetSpecies(speciesID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) PetStats(speciesID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PVP API
func (c *WoWClient) PvpLeaderboards(bracket string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// QUEST API
func (c *WoWClient) Quest(questID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// REALM STATUS API
func (c *WoWClient) RealmStatus() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// RECIPE API
func (c *WoWClient) Recipe(recipeID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// SPELL API
func (c *WoWClient) Spell(spellID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ZONE API
func (c *WoWClient) ZoneMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) Zone(zoneID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// DATA RESOURCES
func (c *WoWClient) Battlegroups() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) CharacterRaces() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) CharacterClasses() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) CharacterAchievements() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) GuildRewards() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) GuildPerks() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) GuildAchievements() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) ItemClasses() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) Talents() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

func (c *WoWClient) PetTypes() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}
