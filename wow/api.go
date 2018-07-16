package wow

/*
	World of Warcraft related API methods should go in here.
*/

// ACHIEVEMENT API
// Achievement provides data about an individual achievement.
func (c *WoWClient) Achievement(id int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// AUCTION API
// Auction APIs currently provide rolling batches of data about current auctions. Fetching auction dumps is
// a two step process that involves checking a per-realm index file to determine if a recent dump has been
// generated and then fetching the most recently generated dump file if necessary.
// AuctionDataStatus provides a per-realm list of recently generated auction house data dumps.
func (c *WoWClient) AuctionDataStatus(realm string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// BOSS API
// BossMasterList lists all supported bosses. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *WoWClient) BossMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// Boss provides information about bosses. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *WoWClient) Boss(bossID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CHALLENGE MODE API
// RealmLeaderboard has data for all 9 challenge mode maps (currently). The map field includes the current medal times for each dungeon.
// Inside each ladder we provide data about each character that was part of each run. The character data includes the current cached spec
// of the character while the member field includes the spec of the character during the challenge mode run.
func (c *WoWClient) RealmLeaderboard(realm string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// RegionLeaderboard has the exact same data format as the realm leaderboards except there is no realm field. It is simply the top 100 results
// gathered for each map for all of the available realm leaderboards in a region.
func (c *WoWClient) RegionLeaderboard() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CHARACTER PROFILE API
// CharacterProfile
func (c *WoWClient) CharacterProfile(characterName string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// GUILD PROFILE API
// GuildProfile
func (c *WoWClient) GuildProfile(guildName string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ITEM API
// Item
func (c *WoWClient) Item(itemID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ItemSet
func (c *WoWClient) ItemSet(steID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// MOUNT API
// MountMasterList
func (c *WoWClient) MountMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PET API
// PetMasterList
func (c *WoWClient) PetMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PetAbilities
func (c *WoWClient) PetAbilities(abilityID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PetSpecies
func (c *WoWClient) PetSpecies(speciesID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PetStats
func (c *WoWClient) PetStats(speciesID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PVP API
// PvpLeaderboards
func (c *WoWClient) PvpLeaderboards(bracket string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// QUEST API
// Quest
func (c *WoWClient) Quest(questID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// REALM STATUS API
// RealmStatus
func (c *WoWClient) RealmStatus() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// RECIPE API
// Recipe
func (c *WoWClient) Recipe(recipeID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// SPELL API
// Spell
func (c *WoWClient) Spell(spellID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ZONE API
// ZoneMasterList
func (c *WoWClient) ZoneMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// Zone
func (c *WoWClient) Zone(zoneID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// DATA RESOURCES
// Battlegroups
func (c *WoWClient) Battlegroups() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CharacterRaces
func (c *WoWClient) CharacterRaces() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CharacterClasses
func (c *WoWClient) CharacterClasses() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CharacterAchievements
func (c *WoWClient) CharacterAchievements() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// GuildRewards
func (c *WoWClient) GuildRewards() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// GuildPerks
func (c *WoWClient) GuildPerks() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// GuildAchievements
func (c *WoWClient) GuildAchievements() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ItemClasses
func (c *WoWClient) ItemClasses() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// Talents
func (c *WoWClient) Talents() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PetTypes
func (c *WoWClient) PetTypes() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}
