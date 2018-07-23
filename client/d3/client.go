package d3

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/models/d3"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

// Client allows the user to access the Diablo III Battle.net API.
type Client struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	region    regions.Region
	key       string
}

// New creates a new Client. Passing different interface values/types
// can cause different behaviors. See function definition for more details.
func New(s *settings.BNetSettings) (c *Client, err error) {
	c = &Client{
		userAgent: "GoBattleNetD3/" + settings.ClientVersion,
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		region:    regions.US,
		key:       "",
	}

	if nil == s {
		return c, nil
	}

	if s.Region.Int() > 5 {
		return nil, errors.ErrUnsupportedArgument
	}

	c.client = s.Client
	c.locale = s.Locale
	c.region = s.Region
	c.key = s.Key

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

// ActIndex gets an index of acts.
func (c *Client) ActIndex() (*d3.ActIndex, error) {
	var index *d3.ActIndex

	err := c.get(endpointActs(c.region), &index)

	if nil != err {
		return nil, err
	}

	return index, nil
}

// Act gets a single act by id.
func (c *Client) Act(actID int) (*d3.Act, error) {
	var act *d3.Act

	err := c.get(endpointAct(c.region, actID), &act)

	if nil != err {
		return nil, err
	}

	return act, nil
}

// Artisan gets a single artisan by slug.
func (c *Client) Artisan(artisanSlug string) (*d3.Artisan, error) {
	var artisan *d3.Artisan

	err := c.get(endpointArtisan(c.region, artisanSlug), &artisan)

	if nil != err {
		return nil, err
	}

	return artisan, nil
}

// Recipe gets a single recipe by slug for the specified artisan.
func (c *Client) Recipe(artisanSlug, recipeSlug string) (*d3.Recipe, error) {
	var recipe *d3.Recipe

	err := c.get(endpointRecipe(c.region, artisanSlug, recipeSlug), &recipe)

	if nil != err {
		return nil, err
	}

	return recipe, nil
}

// Follower gets a single follower by slug.
func (c *Client) Follower(followerSlug string) (*d3.Follower, error) {
	var follower *d3.Follower

	err := c.get(endpointFollower(c.region, followerSlug), &follower)

	if nil != err {
		return nil, err
	}

	return follower, nil
}

// CharacterClass gets a single character class by slug.
func (c *Client) CharacterClass(classSlug string) (*d3.CharacterClass, error) {
	var cc *d3.CharacterClass

	err := c.get(endpointCharacterClass(c.region, classSlug), &cc)

	if nil != err {
		return nil, err
	}

	return cc, nil
}

// CharacterSkill gets a single skill by slug, for a specific character class.
func (c *Client) CharacterSkill(classSlug, skillSlug string) (*d3.CharacterAPISkill, error) {
	var cs *d3.CharacterAPISkill

	err := c.get(endpointSkill(c.region, classSlug, skillSlug), &cs)

	if nil != err {
		return nil, err
	}

	return cs, nil
}

// ItemTypeIndex gets an index of item types.
func (c *Client) ItemTypeIndex() (*d3.ItemTypeIndex, error) {
	var index *d3.ItemTypeIndex

	err := c.get(endpointItemTypeIndex(c.region), &index)

	if nil != err {
		return nil, err
	}

	return index, nil
}

// ItemType gets a single item type by slug.
func (c *Client) ItemType(itemTypeSlug string) (*d3.ItemType, error) {
	var it *d3.ItemType

	err := c.get(endpointItemType(c.region, itemTypeSlug), &it)

	if nil != err {
		return nil, err
	}

	return it, nil
}

// Item gets a single item by item slug and ID.
func (c *Client) Item(itemSlugAndID string) (*d3.Item, error) {
	var item *d3.Item

	err := c.get(endpointItem(c.region, itemSlugAndID), &item)

	if nil != err {
		return nil, err
	}

	return item, nil
}

// Account gets a single profile.
func (c *Client) Account(account string) (*d3.Account, error) {
	var acct *d3.Account

	err := c.get(endpointAccount(c.region, account), &acct)

	if nil != err {
		return nil, err
	}

	return acct, nil
}

// Hero gets a single hero.
func (c *Client) Hero(account string, heroID int) (*d3.Hero, error) {
	var hero *d3.Hero

	err := c.get(endpointHero(c.region, account, heroID), &hero)

	if nil != err {
		return nil, err
	}

	return hero, nil
}

// HeroItems gets a list of items for the specified hero.
func (c *Client) HeroItems(account string, heroID int) (*d3.HeroItems, error) {
	var items *d3.HeroItems

	err := c.get(endpointDetailedHeroItems(c.region, account, heroID), &items)

	if nil != err {
		return nil, err
	}

	return items, nil
}

// FollowerItems gets a list of items for the specified hero's followers.
func (c *Client) FollowerItems(account string, heroID int) (*d3.HeroFollowers, error) {
	var followers *d3.HeroFollowers

	err := c.get(endpointDetailedFollowerItems(c.region, account, heroID), &followers)

	if nil != err {
		return nil, err
	}

	return followers, nil
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
