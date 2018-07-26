package battlenet

import (
	"github.com/munsy/gobattlenet/regions"
)

// Regions holds all of the possible Battle.net regions.
var Regions = struct {
	US  regions.Region
	EU  regions.Region
	KR  regions.Region
	TW  regions.Region
	SEA regions.Region
	CN  regions.Region
}{
	US:  regions.US,
	EU:  regions.EU,
	KR:  regions.KR,
	TW:  regions.TW,
	SEA: regions.SEA,
	CN:  regions.CN,
}
