package wow

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/models/wow"
	"github.com/munsy/gobattlenet/pkg/quota"
	"github.com/munsy/gobattlenet/pkg/regions"
)

// Service represents the World of Warcraft service.
type Service struct {
	client *http.Client
	region regions.Region
	locale locale.Locale
	key    string
}

// New returns a new World of Warcraft service.
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

// Achievement provides data about an individual achievement.
func (s *Service) Achievement(id int) (*Response, error) {
	var achievement *wow.Achievement

	return s.get(endpointAchievement(s.region, id), &achievement)
}

// AuctionDataStatus provides a per-realm list of recently generated auction house data dumps.
// Auction APIs currently provide rolling batches of data about current auctions. Fetching auction dumps is
// a two step process that involves checking a per-realm index file to determine if a recent dump has been
// generated and then fetching the most recently generated dump file if necessary.
func (s *Service) AuctionDataStatus(realm string) (*Response, error) {
	var file *wow.File

	_, err := s.get(endpointAuctionDataStatus(s.region, realm), &file)

	if nil != err {
		return nil, err
	}

	var auctionData *wow.AuctionJSONFileData

	return s.get(file.URL, &auctionData)
}

// BossMasterList lists all supported bosses. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (s *Service) BossMasterList() (*Response, error) {
	var bosses *wow.BossMasterList

	return s.get(endpointBossMasterList(s.region), &bosses)
}

// Boss provides information about a boss. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (s *Service) Boss(bossID int) (*Response, error) {
	var boss *wow.Boss

	return s.get(endpointBossInfo(s.region, bossID), &boss)
}

// CharacterProfile is the primary way to access character information. This Character Profile API can be used to fetch a single character at a
// time through an HTTP GET request to a URL describing the character profile resource. By default, a basic dataset will be returned and with each
// request and zero or more additional fields can be retrieved. To access this API, craft a resource URL pointing to the character who's information
// is to be retrieved.
// Optional fields:
// achievements,appearance,feed,guild,hunterPets,items,mounts,pets,petSlots,professions,progression,pvp,quests,reputation,statistics,stats,talents,titles,audit
func (s *Service) CharacterProfile(realm, characterName string) (*Response, error) {
	var data *wow.CharacterData

	return s.get(endpointCharacterProfile(s.region, realm, characterName), &data)
}

// GuildProfile is the primary way to access guild information. This guild profile API can be used to fetch a single guild at a time through an HTTP GET
// request to a url describing the guild profile resource. By default, a basic dataset will be returned and with each request and zero or more additional
// fields can be retrieved.
//
// There are no required query string parameters when accessing this resource, although the fields query string parameter can optionally be passed to indicate
// that one or more of the optional datasets is to be retrieved. Those additional fields are listed in the method titled "Optional Fields".
func (s *Service) GuildProfile(realm, guildName string) (*Response, error) {
	var profile *wow.GuildProfile

	return s.get(endpointGuildProfile(s.region, realm, guildName), &profile)
}

// Item provides detailed item information. This includes item set information if this item is part of a set.
func (s *Service) Item(itemID int) (*Response, error) {
	var item *wow.Item

	return s.get(endpointItemID(s.region, itemID), &item)
}

// ItemSet provides detailed item information. This includes item set information if this item is part of a set.
func (s *Service) ItemSet(setID int) (*Response, error) {
	var set *wow.ItemSet

	return s.get(endpointItemSet(s.region, setID), &set)
}

// MountMasterList returns a list of all supported mounts.
func (s *Service) MountMasterList() (*Response, error) {
	var list *wow.MountList

	return s.get(endpointMount(s.region), &list)
}

// PetMasterList returns a list of all supported battle and vanity pets.
func (s *Service) PetMasterList() (*Response, error) {
	var data *wow.PetData

	return s.get(endpointPetMasterList(s.region), &data)
}

// PetAbilities provides data about a individual battle pet ability ID. We do not provide the tooltip for the ability
// yet. We are working on a better way to provide this since it depends on your pet's species, level and quality rolls.
func (s *Service) PetAbilities(abilityID int) (*Response, error) {
	var ability *wow.PetAbility

	return s.get(endpointPetAbilities(s.region, abilityID), &ability)
}

// PetSpecies provides the data about an individual pet species. The species IDs can be found your character profile using
// the options pets field. Each species also has data about what it's 6 abilities are.
func (s *Service) PetSpecies(speciesID int) (*Response, error) {
	var species *wow.PetSpecies

	return s.get(endpointPetSpecies(s.region, speciesID), &species)
}

// PetStats retrieves detailed information about the given species of pet.
func (s *Service) PetStats(speciesID int) (*Response, error) {
	var stats *wow.PetStats

	return s.get(endpointPetStats(s.region, speciesID), &stats)
}

// PvpLeaderboards provides leaderboard information for the 2v2, 3v3, 5v5 and Rated Battleground leaderboards.
/* Disabled until BFA.
func (s *Service) PvpLeaderboards(bracket string) (*Response, error) {
	var thing *wow.Thing

	return s.get(thing(s.region), thing)
}
*/

// Quest retrieves metadata for a given quest.
func (s *Service) Quest(questID int) (*Response, error) {
	var quest *wow.Quest

	return s.get(endpointQuestID(s.region, questID), &quest)
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
func (s *Service) RealmStatus() (*Response, error) {
	var status *wow.RealmStatus

	return s.get(endpointRealmStatus(s.region), &status)
}

// Recipe provides basic recipe information.
func (s *Service) Recipe(recipeID int) (*Response, error) {
	var recipe *wow.Recipe

	return s.get(endpointRecipeID(s.region, recipeID), &recipe)
}

// Spell provides some information about spells.
func (s *Service) Spell(spellID int) (*Response, error) {
	var spell *wow.Spell

	return s.get(endpointSpellID(s.region, spellID), &spell)
}

// ZoneMasterList returns a list of all supported zones and their bosses. A 'zone' in this context should be considered a dungeon, or a
// raid, not a zone as in a world zone. A 'boss' in this context should be considered a boss encounter, which may include more than one NPC.
func (s *Service) ZoneMasterList() (*Response, error) {
	var list *wow.ZoneMasterList

	return s.get(endpointZoneMasterList(s.region), &list)
}

// Zone provides some information about zones.
func (s *Service) Zone(zoneID int) (*Response, error) {
	var zone *wow.Zone

	return s.get(endpointZone(s.region, zoneID), &zone)
}

// Battlegroups provides the list of battlegroups for this region
func (s *Service) Battlegroups() (*Response, error) {
	var battlegroup *wow.BattleGroupsData

	return s.get(endpointBattlegroups(s.region), &battlegroup)
}

// CharacterRaces provides a list of each race and their associated faction, name, unique ID, and skin.
func (s *Service) CharacterRaces() (*Response, error) {
	var races *wow.CharacterRacesData

	return s.get(endpointCharacterRaces(s.region), &races)
}

// CharacterClasses provides a list of character classes.
func (s *Service) CharacterClasses() (*Response, error) {
	var classes *wow.CharacterClassesData

	return s.get(endpointCharacterClasses(s.region), &classes)
}

// CharacterAchievements provides a list of all of the achievements that characters can earn as well as the category structure and hierarchy.
func (s *Service) CharacterAchievements() (*Response, error) {
	var data *wow.CharacterAchievementsData

	return s.get(endpointCharacterAchievements(s.region), &data)
}

// GuildRewards provides a list of all guild rewards.
func (s *Service) GuildRewards() (*Response, error) {
	var data *wow.GuildRewardsData

	return s.get(endpointGuildRewards(s.region), &data)
}

// GuildPerks provides a list of all guild perks.
func (s *Service) GuildPerks() (*Response, error) {
	var data *wow.GuildPerksData

	return s.get(endpointGuildPerks(s.region), &data)
}

// GuildAchievements provides a list of all of the achievements that guilds can earn as well as the category structure and hierarchy.
func (s *Service) GuildAchievements() (*Response, error) {
	var data *wow.GuildAchievementsData

	return s.get(endpointGuildAchievements(s.region), &data)
}

// ItemClasses provides a list of item classes.
func (s *Service) ItemClasses() (*Response, error) {
	var data *wow.ItemClassesData

	return s.get(endpointItemClasses(s.region), &data)
}

// Talents provides a list of talents, specs and glyphs for each class.
func (s *Service) Talents() (*Response, error) {
	var data *wow.TalentsData

	return s.get(endpointTalents(s.region), &data)
}

// PetTypes provides a list of different pet types (including what they are strong and weak against).
func (s *Service) PetTypes() (*Response, error) {
	var data *wow.PetTypesData

	return s.get(endpointPetTypes(s.region), &data)
}
