package response

import (
	"github.com/munsy/battlenet/models/sc2"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// LadderSeasons represents a LadderSeasons response.
type LadderSeasons struct {
	Data     *sc2.LadderSeasons
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}
