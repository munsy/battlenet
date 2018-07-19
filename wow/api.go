package wow

/*
	World of Warcraft related API methods should go in here.
*/

// ACHIEVEMENT API
// Achievement provides data about an individual achievement.
func (c *WoWClient) Achievement(id int) (*Achievement, error) {
	var achievement *Achievement

	err := c.get(endpointAchievement(c.region, id), achievement)

	if nil != err {
		return nil, err
	}

	return achievement, nil
}

// AUCTION API
// Auction APIs currently provide rolling batches of data about current auctions. Fetching auction dumps is
// a two step process that involves checking a per-realm index file to determine if a recent dump has been
// generated and then fetching the most recently generated dump file if necessary.
// AuctionDataStatus provides a per-realm list of recently generated auction house data dumps.
func (c *WoWClient) AuctionDataStatus(realm string) (*AuctionJSONFileData, error) {
	var file *File

	err := c.get(endpointAuctionDataStatus(c.region, realm), file)

	if nil != err {
		return nil, err
	}

	var auctionData *AuctionJSONFileData

	err = c.get(file.URL, auctionData)

	if nil != err {
		return nil, err
	}

	return auctionData, nil
}

// BOSS API
// BossMasterList lists all supported bosses. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *WoWClient) BossMasterList() (*BossMasterList, error) {
	var bosses *BossMasterList

	err := c.get(endpointBossMasterList(c.region), bosses)

	if nil != err {
		return nil, err
	}

	return bosses, nil
}

// Boss provides information about a boss. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *WoWClient) Boss(bossID int) (*Boss, error) {
	var boss *Boss

	err := c.get(endpointBossInfo(c.region, bossID), boss)

	if nil != err {
		return nil, err
	}

	return boss, nil
}

// CHARACTER PROFILE API
// CharacterProfile is the primary way to access character information. This Character Profile API can be used to fetch a single character at a
// time through an HTTP GET request to a URL describing the character profile resource. By default, a basic dataset will be returned and with each
// request and zero or more additional fields can be retrieved. To access this API, craft a resource URL pointing to the character who's information
// is to be retrieved.
func (c *WoWClient) CharacterProfile(characterName string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// GUILD PROFILE API
// GuildProfile is the primary way to access guild information. This guild profile API can be used to fetch a single guild at a time through an HTTP GET
// request to a url describing the guild profile resource. By default, a basic dataset will be returned and with each request and zero or more additional
// fields can be retrieved.
//
// There are no required query string parameters when accessing this resource, although the fields query string parameter can optionally be passed to indicate
// that one or more of the optional datasets is to be retrieved. Those additional fields are listed in the method titled "Optional Fields".
func (c *WoWClient) GuildProfile(guildName string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ITEM API
// Item provides detailed item information. This includes item set information if this item is part of a set.
func (c *WoWClient) Item(itemID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ItemSet provides detailed item information. This includes item set information if this item is part of a set.
func (c *WoWClient) ItemSet(steID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// MOUNT API
// MountMasterList returns a list of all supported mounts.
func (c *WoWClient) MountMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PET API
// PetMasterList returns a list of all supported battle and vanity pets.
func (c *WoWClient) PetMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PetAbilities provides data about a individual battle pet ability ID. We do not provide the tooltip for the ability
// yet. We are working on a better way to provide this since it depends on your pet's species, level and quality rolls.
func (c *WoWClient) PetAbilities(abilityID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PetSpecies provides the data about an individual pet species. The species IDs can be found your character profile using
// the options pets field. Each species also has data about what it's 6 abilities are.
func (c *WoWClient) PetSpecies(speciesID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PetStats retrieves detailed information about the given species of pet.
func (c *WoWClient) PetStats(speciesID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PVP API
// PvpLeaderboards provides leaderboard information for the 2v2, 3v3, 5v5 and Rated Battleground leaderboards.
func (c *WoWClient) PvpLeaderboards(bracket string) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// QUEST API
// Quest retrieves metadata for a given quest.
func (c *WoWClient) Quest(questID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// REALM STATUS API
// RealmStatus allows developers to retrieve realm status information. This information is limited to whether or not the
// realm is up, the type and state of the realm, the current population, and the status of the two world PvP zones.
//
// There are no required query string parameters when accessing this resource, although the optional realms parameter can
// be used to limit the realms returned to a specific set of realms.
//
// PvP Area Status Fields
//
// area - An internal id of this zone.
// controlling-faction - Which faction is controlling the zone at the moment. Possible values are: 0) Alliance, 1) Horde, 2) Neutral
// status - The current status of the zone. The possible values are: -1) Unknown, 0) Idle, 1) Populating, 2) Active, 3) Concluded,
// next - A timestamp of when the next battle starts.
func (c *WoWClient) RealmStatus() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// RECIPE API
// Recipe provides basic recipe information.
func (c *WoWClient) Recipe(recipeID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// SPELL API
// Spell provides some information about spells.
func (c *WoWClient) Spell(spellID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ZONE API
// ZoneMasterList returns a list of all supported zones and their bosses. A 'zone' in this context should be considered a dungeon, or a
// raid, not a zone as in a world zone. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *WoWClient) ZoneMasterList() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// Zone provides some information about zones.
func (c *WoWClient) Zone(zoneID int) (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// DATA RESOURCES
// Battlegroups provides the list of battlegroups for this region
func (c *WoWClient) Battlegroups() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CharacterRaces provides a list of each race and their associated faction, name, unique ID, and skin.
func (c *WoWClient) CharacterRaces() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CharacterClasses provides a list of character classes.
func (c *WoWClient) CharacterClasses() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// CharacterAchievements provides a list of all of the achievements that characters can earn as well as the category structure and hierarchy.
func (c *WoWClient) CharacterAchievements() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// GuildRewards provides a list of all guild rewards.
func (c *WoWClient) GuildRewards() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// GuildPerks provides a list of all guild perks.
func (c *WoWClient) GuildPerks() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// GuildAchievements provides a list of all of the achievements that guilds can earn as well as the category structure and hierarchy.
func (c *WoWClient) GuildAchievements() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// ItemClasses provides a list of item classes.
func (c *WoWClient) ItemClasses() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// Talents provides a list of talents, specs and glyphs for each class.
func (c *WoWClient) Talents() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}

// PetTypes provides a list of different pet types (including what they are strong and weak against).
func (c *WoWClient) PetTypes() (*Thing, error) {
	var thing *Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}
