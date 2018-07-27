package response

import (
	"github.com/munsy/battlenet/models/wow"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// WoWCharacters represents a WoWCharacters response.
type WoWCharacters struct {
	Data     *wow.Characters
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}
