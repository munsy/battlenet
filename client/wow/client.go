package wow

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/models/wow"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

// Client allows the user to access the World of Warcraft Battle.net API.
type Client struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	region    regions.Region
	key       string
}

// New creates a new Client. Passing different interface values/types
// can cause different behaviors. See function definition for more details.
func New(args ...interface{}) (c *Client, err error) {
	c = &Client{
		userAgent: "GoBattleNetWow/" + settings.ClientVersion,
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		region:    regions.US,
		key:       "",
	}

	if nil == args {
		return c, nil
	}

	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			c.key = t
			break
		case *http.Client:
			c.client = t
			break
		case locale.Locale:
			c.locale = t
			break
		case regions.Region:
			c.region = t
			break
		case settings.BNetSettings:
			c.client = t.Client
			c.locale = t.Locale
			c.region = t.Region
			c.key = t.Key
			break
		default:
			return nil, errors.ErrUnsupportedArgument
		}
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

// Achievement provides data about an individual achievement.
func (c *Client) Achievement(id int) (*wow.Achievement, error) {
	var achievement *wow.Achievement

	err := c.get(endpointAchievement(c.region, id), &achievement)

	if nil != err {
		return nil, err
	}

	return achievement, nil
}

// AuctionDataStatus provides a per-realm list of recently generated auction house data dumps.
// Auction APIs currently provide rolling batches of data about current auctions. Fetching auction dumps is
// a two step process that involves checking a per-realm index file to determine if a recent dump has been
// generated and then fetching the most recently generated dump file if necessary.
func (c *Client) AuctionDataStatus(realm string) (*wow.AuctionJSONFileData, error) {
	var file *wow.File

	err := c.get(endpointAuctionDataStatus(c.region, realm), &file)

	if nil != err {
		return nil, err
	}

	var auctionData *wow.AuctionJSONFileData

	err = c.get(file.URL, &auctionData)

	if nil != err {
		return nil, err
	}

	return auctionData, nil
}

// BossMasterList lists all supported bosses. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *Client) BossMasterList() (*wow.BossMasterList, error) {
	var bosses *wow.BossMasterList

	err := c.get(endpointBossMasterList(c.region), &bosses)

	if nil != err {
		return nil, err
	}

	return bosses, nil
}

// Boss provides information about a boss. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *Client) Boss(bossID int) (*wow.Boss, error) {
	var boss *wow.Boss

	err := c.get(endpointBossInfo(c.region, bossID), &boss)

	if nil != err {
		return nil, err
	}

	return boss, nil
}

// CharacterProfile is the primary way to access character information. This Character Profile API can be used to fetch a single character at a
// time through an HTTP GET request to a URL describing the character profile resource. By default, a basic dataset will be returned and with each
// request and zero or more additional fields can be retrieved. To access this API, craft a resource URL pointing to the character who's information
// is to be retrieved.
// Optional fields:
// achievements,appearance,feed,guild,hunterPets,items,mounts,pets,petSlots,professions,progression,pvp,quests,reputation,statistics,stats,talents,titles,audit
func (c *Client) CharacterProfile(realm, characterName string) (*wow.CharacterData, error) {
	var data *wow.CharacterData

	err := c.get(endpointCharacterProfile(c.region, realm, characterName), &data)

	if nil != err {
		return nil, err
	}

	return data, nil
}

// GuildProfile is the primary way to access guild information. This guild profile API can be used to fetch a single guild at a time through an HTTP GET
// request to a url describing the guild profile resource. By default, a basic dataset will be returned and with each request and zero or more additional
// fields can be retrieved.
//
// There are no required query string parameters when accessing this resource, although the fields query string parameter can optionally be passed to indicate
// that one or more of the optional datasets is to be retrieved. Those additional fields are listed in the method titled "Optional Fields".
func (c *Client) GuildProfile(realm, guildName string) (*wow.GuildProfile, error) {
	var profile *wow.GuildProfile

	err := c.get(endpointGuildProfile(c.region, realm, guildName), &profile)

	if nil != err {
		return nil, err
	}

	return profile, nil
}

// Item provides detailed item information. This includes item set information if this item is part of a set.
func (c *Client) Item(itemID int) (*wow.Item, error) {
	var item *wow.Item

	err := c.get(endpointItemID(c.region, itemID), &item)

	if nil != err {
		return nil, err
	}

	return item, nil
}

// ItemSet provides detailed item information. This includes item set information if this item is part of a set.
func (c *Client) ItemSet(setID int) (*wow.ItemSet, error) {
	var set *wow.ItemSet

	err := c.get(endpointItemSet(c.region, setID), &set)

	if nil != err {
		return nil, err
	}

	return set, nil
}

// MountMasterList returns a list of all supported mounts.
func (c *Client) MountMasterList() (*wow.MountList, error) {
	var list *wow.MountList

	err := c.get(endpointMount(c.region), &list)

	if nil != err {
		return nil, err
	}

	return list, nil
}

// PetMasterList returns a list of all supported battle and vanity pets.
func (c *Client) PetMasterList() (*wow.PetData, error) {
	var data *wow.PetData

	err := c.get(endpointPetMasterList(c.region), &data)

	if nil != err {
		return nil, err
	}

	return data, nil
}

// PetAbilities provides data about a individual battle pet ability ID. We do not provide the tooltip for the ability
// yet. We are working on a better way to provide this since it depends on your pet's species, level and quality rolls.
func (c *Client) PetAbilities(abilityID int) (*wow.PetAbility, error) {
	var ability *wow.PetAbility

	err := c.get(endpointPetAbilities(c.region, abilityID), &ability)

	if nil != err {
		return nil, err
	}

	return ability, nil
}

// PetSpecies provides the data about an individual pet species. The species IDs can be found your character profile using
// the options pets field. Each species also has data about what it's 6 abilities are.
func (c *Client) PetSpecies(speciesID int) (*wow.PetSpecies, error) {
	var species *wow.PetSpecies

	err := c.get(endpointPetSpecies(c.region, speciesID), &species)

	if nil != err {
		return nil, err
	}

	return species, nil
}

// PetStats retrieves detailed information about the given species of pet.
func (c *Client) PetStats(speciesID int) (*wow.PetStats, error) {
	var stats *wow.PetStats

	err := c.get(endpointPetStats(c.region, speciesID), &stats)

	if nil != err {
		return nil, err
	}

	return stats, nil
}

// PvpLeaderboards provides leaderboard information for the 2v2, 3v3, 5v5 and Rated Battleground leaderboards.
/* Disabled until BFA.
func (c *Client) PvpLeaderboards(bracket string) (*wow.Thing, error) {
	var thing *wow.Thing

	err := c.get(thing(c.region), thing)

	if nil != err {
		return nil, err
	}

	return thing, nil
}
*/

// Quest retrieves metadata for a given quest.
func (c *Client) Quest(questID int) (*wow.Quest, error) {
	var quest *wow.Quest

	err := c.get(endpointQuestID(c.region, questID), &quest)

	if nil != err {
		return nil, err
	}

	return quest, nil
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
func (c *Client) RealmStatus() (*wow.RealmStatus, error) {
	var status *wow.RealmStatus

	err := c.get(endpointRealmStatus(c.region), &status)

	if nil != err {
		return nil, err
	}

	return status, nil
}

// Recipe provides basic recipe information.
func (c *Client) Recipe(recipeID int) (*wow.Recipe, error) {
	var recipe *wow.Recipe

	err := c.get(endpointRecipeID(c.region, recipeID), &recipe)

	if nil != err {
		return nil, err
	}

	return recipe, nil
}

// Spell provides some information about spells.
func (c *Client) Spell(spellID int) (*wow.Spell, error) {
	var spell *wow.Spell

	err := c.get(endpointSpellID(c.region, spellID), &spell)

	if nil != err {
		return nil, err
	}

	return spell, nil
}

// ZoneMasterList returns a list of all supported zones and their bosses. A 'zone' in this context should be considered a dungeon, or a
// raid, not a zone as in a world zone. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (c *Client) ZoneMasterList() (*wow.ZoneMasterList, error) {
	var list *wow.ZoneMasterList

	err := c.get(endpointZoneMasterList(c.region), &list)

	if nil != err {
		return nil, err
	}

	return list, nil
}

// Zone provides some information about zones.
func (c *Client) Zone(zoneID int) (*wow.Zone, error) {
	var zone *wow.Zone

	err := c.get(endpointZone(c.region, zoneID), &zone)

	if nil != err {
		return nil, err
	}

	return zone, nil
}

// Battlegroups provides the list of battlegroups for this region
func (c *Client) Battlegroups() (*wow.BattleGroupsData, error) {
	var battlegroup *wow.BattleGroupsData

	err := c.get(endpointBattlegroups(c.region), &battlegroup)

	if nil != err {
		return nil, err
	}

	return battlegroup, nil
}

// CharacterRaces provides a list of each race and their associated faction, name, unique ID, and skin.
func (c *Client) CharacterRaces() (*wow.CharacterRacesData, error) {
	var races *wow.CharacterRacesData

	err := c.get(endpointCharacterRaces(c.region), &races)

	if nil != err {
		return nil, err
	}

	return races, nil
}

// CharacterClasses provides a list of character classes.
func (c *Client) CharacterClasses() (*wow.CharacterClassesData, error) {
	var classes *wow.CharacterClassesData

	err := c.get(endpointCharacterClasses(c.region), &classes)

	if nil != err {
		return nil, err
	}

	return classes, nil
}

// CharacterAchievements provides a list of all of the achievements that characters can earn as well as the category structure and hierarchy.
func (c *Client) CharacterAchievements() (*wow.CharacterAchievementsData, error) {
	var data *wow.CharacterAchievementsData

	err := c.get(endpointCharacterAchievements(c.region), &data)

	if nil != err {
		return nil, err
	}

	return data, nil
}

// GuildRewards provides a list of all guild rewards.
func (c *Client) GuildRewards() (*wow.GuildRewardsData, error) {
	var data *wow.GuildRewardsData

	err := c.get(endpointGuildRewards(c.region), &data)

	if nil != err {
		return nil, err
	}

	return data, nil
}

// GuildPerks provides a list of all guild perks.
func (c *Client) GuildPerks() (*wow.GuildPerksData, error) {
	var data *wow.GuildPerksData

	err := c.get(endpointGuildPerks(c.region), &data)

	if nil != err {
		return nil, err
	}

	return data, nil
}

// GuildAchievements provides a list of all of the achievements that guilds can earn as well as the category structure and hierarchy.
func (c *Client) GuildAchievements() (*wow.GuildAchievementsData, error) {
	var data *wow.GuildAchievementsData

	err := c.get(endpointGuildAchievements(c.region), &data)

	if nil != err {
		return nil, err
	}

	return data, nil
}

// ItemClasses provides a list of item classes.
func (c *Client) ItemClasses() (*wow.ItemClassesData, error) {
	var data *wow.ItemClassesData

	err := c.get(endpointItemClasses(c.region), &data)

	if nil != err {
		return nil, err
	}

	return data, nil
}

// Talents provides a list of talents, specs and glyphs for each class.
func (c *Client) Talents() (*wow.TalentsData, error) {
	var data *wow.TalentsData

	err := c.get(endpointTalents(c.region), &data)

	if nil != err {
		return nil, err
	}

	return data, nil
}

// PetTypes provides a list of different pet types (including what they are strong and weak against).
func (c *Client) PetTypes() (*wow.PetTypesData, error) {
	var data *wow.PetTypesData

	err := c.get(endpointPetTypes(c.region), &data)

	if nil != err {
		return nil, err
	}

	return data, nil
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
