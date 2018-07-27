package response

import (
	"github.com/munsy/battlenet/models/account"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// BattleID represents a BattleID response.
type BattleID struct {
	Data     *account.BattleID
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}
