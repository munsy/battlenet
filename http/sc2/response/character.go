package response

import (
	"github.com/munsy/battlenet/models/sc2"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// Character represents a Character response.
type Character struct {
	Data     *sc2.Character
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}
