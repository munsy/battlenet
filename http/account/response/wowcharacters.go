package response

import (
	"github.com/munsy/gobattlenet/models/wow"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

// WoWCharacters represents a WoWCharacters response.
type WoWCharacters struct {
	Data     *wow.Characters
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}
