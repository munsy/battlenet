package account

// BattleID provides the battletag and ID number for a Battle.net account.
type BattleID struct {
	ID        int    `json:"id"`
	BattleTag string `"json:battletag"`
}
