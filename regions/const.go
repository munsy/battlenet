package regions

// Valid regions.
const (
	US Region = iota
	EU
	KR
	TW
	SEA
	CN
)

const urlHead = "https://"
const urlTail = ".api.battle.net/"
const urlCN = "https://api.battlenet.com.cn/"
const oauthTail = ".battle.net/oauth/"
const oauthCN = "https://www.battlenet.com.cn/oauth/"
