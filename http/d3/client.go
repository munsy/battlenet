package d3

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/errors"
	"github.com/munsy/gobattlenet/http/d3/response"
	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/models/d3"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

// Client represents the Diablo III client.
type Client struct {
	client    *http.Client
	region    regions.Region
	locale    locale.Locale
	userAgent string
	key       string
}

// New returns a new Diablo III client.
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

// ActIndex gets an index of acts.
func (c *Client) ActIndex() (*response.ActIndex, error) {
	var data *d3.ActIndex

	ep := endpointActs(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.ActIndex{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Act gets a single act by id.
func (c *Client) Act(actID int) (*response.Act, error) {
	var data *d3.Act

	ep := endpointAct(c.region, actID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Act{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Artisan gets a single artisan by slug.
func (c *Client) Artisan(artisanSlug string) (*response.Artisan, error) {
	var data *d3.Artisan

	ep := endpointArtisan(c.region, artisanSlug)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Artisan{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Recipe gets a single recipe by slug for the specified artisan.
func (c *Client) Recipe(artisanSlug, recipeSlug string) (*response.Recipe, error) {
	var data *d3.Recipe

	ep := endpointRecipe(c.region, artisanSlug, recipeSlug)

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

// Follower gets a single follower by slug.
func (c *Client) Follower(followerSlug string) (*response.Follower, error) {
	var data *d3.Follower

	ep := endpointFollower(c.region, followerSlug)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Follower{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// CharacterClass gets a single character class by slug.
func (c *Client) CharacterClass(classSlug string) (*response.CharacterClass, error) {
	var data *d3.CharacterClass

	ep := endpointCharacterClass(c.region, classSlug)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.CharacterClass{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// CharacterSkill gets a single skill by slug, for a specific character class.
func (c *Client) CharacterSkill(classSlug, skillSlug string) (*response.CharacterSkill, error) {
	var data *d3.CharacterAPISkill

	ep := endpointSkill(c.region, classSlug, skillSlug)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.CharacterSkill{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// ItemTypeIndex gets an index of item types.
func (c *Client) ItemTypeIndex() (*response.ItemTypeIndex, error) {
	var data *d3.ItemTypeIndex

	ep := endpointItemTypeIndex(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.ItemTypeIndex{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// ItemType gets a single item type by slug.
func (c *Client) ItemType(itemTypeSlug string) (*response.ItemType, error) {
	var data *d3.ItemType

	ep := endpointItemType(c.region, itemTypeSlug)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.ItemType{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Item gets a single item by item slug and ID.
func (c *Client) Item(itemSlugAndID string) (*response.Item, error) {
	var data *d3.Item

	ep := endpointItem(c.region, itemSlugAndID)

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

// Account gets a single profile.
func (c *Client) Account(account string) (*response.Account, error) {
	var data *d3.Account

	ep := endpointAccount(c.region, account)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Account{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Hero gets a single hero.
func (c *Client) Hero(account string, heroID int) (*response.Hero, error) {
	var data *d3.Hero

	ep := endpointHero(c.region, account, heroID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Hero{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// HeroItems gets a list of items for the specified hero.
func (c *Client) HeroItems(account string, heroID int) (*response.HeroItems, error) {
	var data *d3.HeroItems

	ep := endpointDetailedHeroItems(c.region, account, heroID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.HeroItems{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// FollowerItems gets a list of items for the specified hero's followers.
func (c *Client) FollowerItems(account string, heroID int) (*response.HeroFollowers, error) {
	var data *d3.HeroFollowers

	ep := endpointDetailedFollowerItems(c.region, account, heroID)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.HeroFollowers{
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
