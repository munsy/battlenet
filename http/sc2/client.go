package sc2

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/errors"
	"github.com/munsy/gobattlenet/http/sc2/response"
	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/models/sc2"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

// Client represents the Starcraft II client.
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

// New returns a new Starcraft II client.
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

// Character returns the Sc2 character profile.
func (c *Client) Character(name string, profileID int) (*response.Character, error) {
	var data *sc2.Character

	ep := endpointProfile(c.region, profileID, c.region.Itoa(), name)

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

// LadderSeasons returns the Sc2 profile's ladder seasons.
func (c *Client) LadderSeasons(name string, profileID int) (*response.LadderSeasons, error) {
	var data *sc2.LadderSeasons

	ep := endpointLadderProfile(c.region, profileID, c.region.Itoa(), name)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.LadderSeasons{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// MatchHistory returns the Sc2 profile's match history.
func (c *Client) MatchHistory(name string, profileID int) (*response.MatchHistory, error) {
	var data *sc2.MatchHistory

	ep := endpointMatchHistory(c.region, profileID, c.region.Itoa(), name)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.MatchHistory{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Ladder returns Sc2 ladder data.
func (c *Client) Ladder(id int) (*response.Ladder, error) {
	var data *sc2.Ladder

	ep := endpointLadder(c.region, id)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Ladder{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Achievements returns Sc2 achievement data.
func (c *Client) Achievements(id int) (*response.Achievements, error) {
	var data *sc2.AchievementsData

	ep := endpointAchievements(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Achievements{
		Data:     data,
		Endpoint: ep,
		Quota:    q,
		Region:   c.region,
	}, nil
}

// Rewards returns Sc2 reward data.
func (c *Client) Rewards(id int) (*response.Rewards, error) {
	var data *sc2.RewardsData

	ep := endpointRewards(c.region)

	q, err := c.get(ep, &data)

	if nil != err {
		return nil, err
	}

	return &response.Rewards{
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
