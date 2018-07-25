package d3

import (
	"encoding/json"

	"github.com/munsy/gobattlenet/pkg/errors"
	"github.com/munsy/gobattlenet/pkg/models/account"
	"github.com/munsy/gobattlenet/pkg/models/sc2"
	"github.com/munsy/gobattlenet/pkg/models/wow"
	"github.com/munsy/gobattlenet/pkg/quota"
	"github.com/munsy/gobattlenet/pkg/regions"
)

// Response represents a Battle.net Account API response.
type Response struct {
	body     []byte
	endpoint string
	quota    *quota.Quota
	region   regions.Region
}

// Data returns the response data as the appropriate type.
func (r *Response) Data() (interface{}, error) {
	var v interface{}

	switch r.endpoint {
	case endpointUser(r.region):
		v = &account.BattleID{}
		break
	case endpointSc2User(r.region):
		v = &sc2.Character{}
		break
	case endpointWowCharacters(r.region):
		v = &wow.Characters{}
		break
	default:
		return nil, errors.ErrUnresolvedEndpoint
	}

	err := json.Unmarshal(r.body, &v)

	if nil != err {
		return nil, err
	}
	return v, nil
}

// Bytes returns the byte slice from the response.
func (r *Response) Bytes() []byte {
	return r.body
}

// Endpoint returns the endpoint used to obtain the response.
func (r *Response) Endpoint() string {
	return r.endpoint
}

// Quota returns details about current Battle.net API quota usage.
func (r Response) Quota() *quota.Quota {
	return r.quota
}
