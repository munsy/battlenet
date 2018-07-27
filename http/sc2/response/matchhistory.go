package response

import (
	"github.com/munsy/battlenet/models/sc2"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// MatchHistory represents a MatchHistory response.
type MatchHistory struct {
	Data     *sc2.MatchHistory
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}
