package response

import (
	"github.com/munsy/gobattlenet/models/wow"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

type WoWCharacters struct {
	Data     *wow.Characters
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}
