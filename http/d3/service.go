package d3

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/models/d3"
	"github.com/munsy/gobattlenet/pkg/quota"
	"github.com/munsy/gobattlenet/pkg/regions"
)

// Service represents the Diablo III service.
type Service struct {
	client *http.Client
	region regions.Region
	locale locale.Locale
	key    string
}

// New returns a new Diablo III service.
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

// ActIndex gets an index of acts.
func (s *Service) ActIndex() (*Response, error) {
	var index *d3.ActIndex

	return s.get(endpointActs(s.region), &index)
}

// Act gets a single act by id.
func (s *Service) Act(actID int) (*Response, error) {
	var act *d3.Act

	return s.get(endpointAct(s.region, actID), &act)
}

// Artisan gets a single artisan by slug.
func (s *Service) Artisan(artisanSlug string) (*Response, error) {
	var artisan *d3.Artisan

	return s.get(endpointArtisan(s.region, artisanSlug), &artisan)
}

// Recipe gets a single recipe by slug for the specified artisan.
func (s *Service) Recipe(artisanSlug, recipeSlug string) (*Response, error) {
	var recipe *d3.Recipe

	return s.get(endpointRecipe(s.region, artisanSlug, recipeSlug), &recipe)
}

// Follower gets a single follower by slug.
func (s *Service) Follower(followerSlug string) (*Response, error) {
	var follower *d3.Follower

	return s.get(endpointFollower(s.region, followerSlug), &follower)
}

// CharacterClass gets a single character class by slug.
func (s *Service) CharacterClass(classSlug string) (*Response, error) {
	var cc *d3.CharacterClass

	return s.get(endpointCharacterClass(s.region, classSlug), &cc)
}

// CharacterSkill gets a single skill by slug, for a specific character class.
func (s *Service) CharacterSkill(classSlug, skillSlug string) (*Response, error) {
	var cs *d3.CharacterAPISkill

	return s.get(endpointSkill(s.region, classSlug, skillSlug), &cs)
}

// ItemTypeIndex gets an index of item types.
func (s *Service) ItemTypeIndex() (*Response, error) {
	var index *d3.ItemTypeIndex

	return s.get(endpointItemTypeIndex(s.region), &index)
}

// ItemType gets a single item type by slug.
func (s *Service) ItemType(itemTypeSlug string) (*Response, error) {
	var it *d3.ItemType

	return s.get(endpointItemType(s.region, itemTypeSlug), &it)
}

// Item gets a single item by item slug and ID.
func (s *Service) Item(itemSlugAndID string) (*Response, error) {
	var item *d3.Item

	return s.get(endpointItem(s.region, itemSlugAndID), &item)
}

// Account gets a single profile.
func (s *Service) Account(account string) (*Response, error) {
	var acct *d3.Account

	return s.get(endpointAccount(s.region, account), &acct)
}

// Hero gets a single hero.
func (s *Service) Hero(account string, heroID int) (*Response, error) {
	var hero *d3.Hero

	return s.get(endpointHero(s.region, account, heroID), &hero)
}

// HeroItems gets a list of items for the specified hero.
func (s *Service) HeroItems(account string, heroID int) (*Response, error) {
	var items *d3.HeroItems

	return s.get(endpointDetailedHeroItems(s.region, account, heroID), &items)
}

// FollowerItems gets a list of items for the specified hero's followers.
func (s *Service) FollowerItems(account string, heroID int) (*Response, error) {
	var followers *d3.HeroFollowers

	return s.get(endpointDetailedFollowerItems(s.region, account, heroID), &followers)
}
