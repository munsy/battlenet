package account

// User provides the battletag and ID number for a Battle.net account.
type User struct {
	ID        int    `json:"id"`
	BattleTag string `"json:battletag"`
}
