package wow

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/errors"
	"github.com/munsy/gobattlenet/http/wow/response"
	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/models/wow"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

// Client represents the World of Warcraft client.
type Client struct {
	client    *http.Client
	region    regions.Region
	locale    locale.Locale
	userAgent string
	key       string
}

// UserAgent returns the client User-Agent header used in API requests.
func (c *Client) UserAgent() string {
	return c.userAgent
}

// New returns a new World of Warcraft client.
func New(c *http.Client, r regions.Region, l locale.Locale, k, v string) (*Client, error) {
	if "" == k {
		return nil, errors.ErrNoKeySupplied
	}

	if "" == v {
		return nil, errors.ErrNoVersionSupplied
	}

	ac := &Client{
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		region:    regions.US,
		userAgent: "GoBattleNet/" + v,
		key:       k,
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

// Achievement provides data about an individual achievement.
func (c *Client) Achievement(id int) (*response.Achievement, error) {
	var data *wow.Achievement

	ep := endpointAchievement(c.region, id)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Achievement{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// AuctionDataStatus provides a per-realm list of recently generated auction house data dumps.
// Auction APIs currently provide rolling batches of data about current auctions. Fetching auction dumps is
// a two step process that involves checking a per-realm index file to determine if a recent dump has been
// generated and then fetching the most recently generated dump file if necessary.
func (c *Client) AuctionDataStatus(realm string) (*response.AuctionJSONFile, error) {
	var file *wow.File

	ep := endpointAuctionDataStatus(c.region, realm)

	q, err := c.get(ep, &file)

	if nil != err {
		return nil, err
	}

	var data *wow.AuctionJSONFileData

	_, err = c.get(file.URL, &data)

	if nil != err {
		return nil, err
	}

	return &response.AuctionJSONFile{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// BossMasterList lists all supported bosses. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *Client) BossMasterList() (*response.BossMasterList, error) {
	var data *wow.BossMasterList

	ep := endpointBossMasterList(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.BossMasterList{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Boss provides information about a boss. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *Client) Boss(bossID int) (*response.Boss, error) {
	var data *wow.Boss

	ep := endpointBossInfo(c.region, bossID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Boss{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// CharacterProfile is the primary way to access character information. This Character Profile API can be used to fetch a single character at a
// time through an HTTP GET request to a URL describing the character profile resource. By default, a basic dataset will be returned and with each
// request and zero or more additional fields can be retrieved. To access this API, craft a resource URL pointing to the character who's information
// is to be retrieved.
// Optional fields:
// achievements,appearance,feed,guild,hunterPets,items,mounts,pets,petSlots,professions,progression,pvp,quests,reputation,statistics,stats,talents,titles,audit
func (c *Client) CharacterProfile(realm, characterName string) (*response.Character, error) {
	var data *wow.CharacterData

	ep := endpointCharacterProfile(c.region, realm, characterName)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Character{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// GuildProfile is the primary way to access guild information. This guild profile API can be used to fetch a single guild at a time through an HTTP GET
// request to a url describing the guild profile resource. By default, a basic dataset will be returned and with each request and zero or more additional
// fields can be retrieved.
//
// There are no required query string parameters when accessing this resource, although the fields query string parameter can optionally be passed to indicate
// that one or more of the optional datasets is to be retrieved. Those additional fields are listed in the method titled "Optional Fields".
func (c *Client) GuildProfile(realm, guildName string) (*response.GuildProfile, error) {
	var data *wow.GuildProfile

	ep := endpointGuildProfile(c.region, realm, guildName)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.GuildProfile{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Item provides detailed item information. This includes item set information if this item is part of a set.
func (c *Client) Item(itemID int) (*response.Item, error) {
	var data *wow.Item

	ep := endpointItemID(c.region, itemID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Item{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// ItemSet provides detailed item information. This includes item set information if this item is part of a set.
func (c *Client) ItemSet(setID int) (*response.ItemSet, error) {
	var data *wow.ItemSet

	ep := endpointItemSet(c.region, setID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.ItemSet{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// MountMasterList returns a list of all supported mounts.
func (c *Client) MountMasterList() (*response.MountList, error) {
	var data *wow.MountList

	ep := endpointMount(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.MountList{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// PetMasterList returns a list of all supported battle and vanity pets.
func (c *Client) PetMasterList() (*response.Pet, error) {
	var data *wow.PetData

	ep := endpointPetMasterList(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Pet{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// PetAbilities provides data about a individual battle pet ability ID. We do not provide the tooltip for the ability
// yet. We are working on a better way to provide this since it depends on your pet's species, level and quality rolls.
func (c *Client) PetAbilities(abilityID int) (*response.PetAbility, error) {
	var data *wow.PetAbility

	ep := endpointPetAbilities(c.region, abilityID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.PetAbility{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// PetSpecies provides the data about an individual pet species. The species IDs can be found your character profile using
// the options pets field. Each species also has data about what it's 6 abilities are.
func (c *Client) PetSpecies(speciesID int) (*response.PetSpecies, error) {
	var data *wow.PetSpecies

	ep := endpointPetSpecies(c.region, speciesID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.PetSpecies{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// PetStats retrieves detailed information about the given species of pet.
func (c *Client) PetStats(speciesID int) (*response.PetStats, error) {
	var data *wow.PetStats

	ep := endpointPetStats(c.region, speciesID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.PetStats{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// PvpLeaderboards provides leaderboard information for the 2v2, 3v3, 5v5 and Rated Battleground leaderboards.
/* Disabled until BFA.
func (c *WoWClient) PvpLeaderboards(bracket string) (*response.Thing, error) {
	var data *wow.Thing
	err := c.get(thing(c.region), thing)
	if nil != err {
		return nil, err
	}
	return thing, nil
}
*/

// Quest retrieves metadata for a given quest.
func (c *Client) Quest(questID int) (*response.Quest, error) {
	var data *wow.Quest

	ep := endpointQuestID(c.region, questID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Quest{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

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
func (c *Client) RealmStatus() (*response.RealmStatus, error) {
	var data *wow.RealmStatus

	ep := endpointRealmStatus(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.RealmStatus{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Recipe provides basic recipe information.
func (c *Client) Recipe(recipeID int) (*response.Recipe, error) {
	var data *wow.Recipe

	ep := endpointRecipeID(c.region, recipeID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Recipe{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Spell provides some information about spells.
func (c *Client) Spell(spellID int) (*response.Spell, error) {
	var data *wow.Spell

	ep := endpointSpellID(c.region, spellID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Spell{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// ZoneMasterList returns a list of all supported zones and their bosses. A 'zone' in this context should be considered a dungeon, or a
// raid, not a zone as in a world zone. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *Client) ZoneMasterList() (*response.ZoneMasterList, error) {
	var data *wow.ZoneMasterList

	ep := endpointZoneMasterList(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.ZoneMasterList{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Zone provides some information about zones.
func (c *Client) Zone(zoneID int) (*response.Zone, error) {
	var data *wow.Zone

	ep := endpointZone(c.region, zoneID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Zone{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Battlegroups provides the list of battlegroups for this region
func (c *Client) Battlegroups() (*response.BattleGroups, error) {
	var data *wow.BattleGroupsData

	ep := endpointBattlegroups(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.BattleGroups{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// CharacterRaces provides a list of each race and their associated faction, name, unique ID, and skin.
func (c *Client) CharacterRaces() (*response.CharacterRaces, error) {
	var data *wow.CharacterRacesData

	ep := endpointCharacterRaces(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.CharacterRaces{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// CharacterClasses provides a list of character classes.
func (c *Client) CharacterClasses() (*response.CharacterClasses, error) {
	var data *wow.CharacterClassesData

	ep := endpointCharacterClasses(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.CharacterClasses{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// CharacterAchievements provides a list of all of the achievements that characters can earn as well as the category structure and hierarchy.
func (c *Client) CharacterAchievements() (*response.CharacterAchievements, error) {
	var data *wow.CharacterAchievementsData

	ep := endpointCharacterAchievements(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.CharacterAchievements{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// GuildRewards provides a list of all guild rewards.
func (c *Client) GuildRewards() (*response.GuildRewards, error) {
	var data *wow.GuildRewardsData

	ep := endpointGuildRewards(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.GuildRewards{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// GuildPerks provides a list of all guild perks.
func (c *Client) GuildPerks() (*response.GuildPerks, error) {
	var data *wow.GuildPerksData

	ep := endpointGuildPerks(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.GuildPerks{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// GuildAchievements provides a list of all of the achievements that guilds can earn as well as the category structure and hierarchy.
func (c *Client) GuildAchievements() (*response.GuildAchievements, error) {
	var data *wow.GuildAchievementsData

	ep := endpointGuildAchievements(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.GuildAchievements{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// ItemClasses provides a list of item classes.
func (c *Client) ItemClasses() (*response.ItemClasses, error) {
	var data *wow.ItemClassesData

	ep := endpointItemClasses(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.ItemClasses{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Talents provides a list of talents, specs and glyphs for each class.
func (c *Client) Talents() (*response.Talents, error) {
	var data *wow.TalentsData

	ep := endpointTalents(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Talents{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// PetTypes provides a list of different pet types (including what they are strong and weak against).
func (c *Client) PetTypes() (*response.PetTypes, error) {
	var data *wow.PetTypesData

	ep := endpointPetTypes(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.PetTypes{
		Data:     data,
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

	request := endpoint + "?locale=" + c.locale.String() + "&apikey=" + c.key

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
