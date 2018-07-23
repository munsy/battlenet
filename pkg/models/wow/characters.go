package wow

// CharacterSpec represents the character specilization information
// for a particular WoW character from a given WoW profile.
type CharacterSpec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}

// Character represents the character information from a particular WoW profile.
type Character struct {
	Name              string        `json:"name"`
	Realm             string        `json:"realm"`
	Battlegroup       string        `json:"battlegroup"`
	Class             int           `json:"class"`
	Race              int           `json:"race"`
	Gender            int           `json:"gender"`
	Level             int           `json:"level"`
	AchievementPoints int           `json:"achievementPoints"`
	Thumbnail         string        `json:"thumbnail"`
	Spec              CharacterSpec `json:"spec,omitempty"`
	Guild             string        `json:"guild,omitempty"`
	GuildRealm        string        `json:"guildRealm,omitempty"`
	LastModified      int           `json:"lastModified"`
}

// Characters type represents a Character slice.
type Characters struct {
	CharacterList []Character `json:"characters"`
}

// CharacterGuildEmblem represents a World of Warcraft character guild emblem.
type CharacterGuildEmblem struct {
	Icon              int    `json:"icon"`
	IconColor         string `json:"iconColor"`
	IconColorID       int    `json:"iconColorId"`
	Border            int    `json:"border"`
	BorderColor       string `json:"borderColor"`
	BorderColorID     int    `json:"borderColorId"`
	BackgroundColor   string `json:"backgroundColor"`
	BackgroundColorID int    `json:"backgroundColorId"`
}

// CharacterGuild represents a World of Warcraft character guild.
type CharacterGuild struct {
	Name              string               `json:"name"`
	Realm             string               `json:"realm"`
	Battlegroup       string               `json:"battlegroup"`
	Members           int                  `json:"members"`
	AchievementPoints int                  `json:"achievementPoints"`
	Emblem            CharacterGuildEmblem `json:"emblem"`
}

// CharacterFeedItem represents a World of Warcraft character feed item.
type CharacterFeedItem struct {
	Type           string      `json:"type"`
	Timestamp      int64       `json:"timestamp"`
	ItemID         int         `json:"itemId,omitempty"`
	Context        string      `json:"context,omitempty"`
	BonusLists     []int       `json:"bonusLists,omitempty"`
	Achievement    Achievement `json:"achievement,omitempty"`
	FeatOfStrength bool        `json:"featOfStrength,omitempty"`
	Criteria       Criteria    `json:"criteria,omitempty"`
	Quantity       int         `json:"quantity,omitempty"`
	Name           string      `json:"name,omitempty"`
}

// CharacterItemSlotTooltipParams represents tooltip parameters for a World of Warcraft character item slot.
type CharacterItemSlotTooltipParams struct {
	Gem0              int `json:"gem0"`
	Gem1              int `json:"gem1"`
	Gem2              int `json:"gem2"`
	TransmogItem      int `json:"transmogItem"`
	TimewalkerLevel   int `json:"timewalkerLevel"`
	AzeritePower0     int `json:"azeritePower0"`
	AzeritePower1     int `json:"azeritePower1"`
	AzeritePower2     int `json:"azeritePower2"`
	AzeritePower3     int `json:"azeritePower3"`
	AzeritePowerLevel int `json:"azeritePowerLevel"`
}

// CharacterItemSlotStat represents a World of Warcraft character item slot stat.
type CharacterItemSlotStat struct {
	Stat   int `json:"stat"`
	Amount int `json:"amount"`
}

// CharacterItemSlotAppearance represents a World of Warcraft character item slot appearance.
type CharacterItemSlotAppearance struct {
	ItemID                      int `json:"itemId"`
	ItemAppearanceModID         int `json:"itemAppearanceModId"`
	TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
}

// CharacterWeaponDamage represents a World of Warcraft character weapon damage.
type CharacterWeaponDamage struct {
	Min      int `json:"min"`
	Max      int `json:"max"`
	ExactMin int `json:"exactMin"`
	ExactMax int `json:"exactMax"`
}

// CharacterWeaponInfo represents World of Warcraft character weapon info.
type CharacterWeaponInfo struct {
	Damage      CharacterWeaponDamage `json:"damage"`
	WeaponSpeed float64               `json:"weaponSpeed"`
	Dps         float64               `json:"dps"`
}

// CharacterArtifactTrait represents a World of Warcraft character artifact trait.
type CharacterArtifactTrait struct {
	ID   int `json:"id"`
	Rank int `json:"rank"`
}

// CharacterWeaponRelic represents a World of Warcraft character weapon relic.
type CharacterWeaponRelic struct {
	Socket     int   `json:"socket"`
	ItemID     int   `json:"itemId"`
	Context    int   `json:"context"`
	BonusLists []int `json:"bonusLists"`
}

// CharacterWeaponSlot represents a World of Warcraft character weapon slot.
type CharacterWeaponSlot struct {
	ID                   int                            `json:"id"`
	Name                 string                         `json:"name"`
	Icon                 string                         `json:"icon"`
	Quality              int                            `json:"quality"`
	ItemLevel            int                            `json:"itemLevel"`
	TooltipParams        CharacterItemSlotTooltipParams `json:"tooltipParams"`
	Stats                []CharacterItemSlotStat        `json:"stats"`
	Armor                int                            `json:"armor"`
	WeaponInfo           CharacterWeaponInfo            `json:"weaponInfo"`
	Context              string                         `json:"context"`
	BonusLists           []int                          `json:"bonusLists"`
	DisplayInfoID        int                            `json:"displayInfoId"`
	ArtifactID           int                            `json:"artifactId"`
	ArtifactAppearanceID int                            `json:"artifactAppearanceId"`
	ArtifactTraits       []CharacterArtifactTrait       `json:"artifactTraits"`
	Relics               []CharacterWeaponRelic         `json:"relics"`
	Appearance           CharacterItemSlotAppearance    `json:"appearance"`
}

// CharacterItemSlot represents a World of Warcraft character item slot.
type CharacterItemSlot struct {
	ID                   int                            `json:"id"`
	Name                 string                         `json:"name"`
	Icon                 string                         `json:"icon"`
	Quality              int                            `json:"quality"`
	ItemLevel            int                            `json:"itemLevel"`
	TooltipParams        CharacterItemSlotTooltipParams `json:"tooltipParams"`
	Stats                []CharacterItemSlotStat        `json:"stats"`
	Armor                int                            `json:"armor"`
	Context              string                         `json:"context"`
	BonusLists           []int                          `json:"bonusLists"`
	DisplayInfoID        int                            `json:"displayInfoId"`
	ArtifactID           int                            `json:"artifactId"`
	ArtifactAppearanceID int                            `json:"artifactAppearanceId"`
	ArtifactTraits       []interface{}                  `json:"artifactTraits"`
	Relics               []interface{}                  `json:"relics"`
	Appearance           CharacterItemSlotAppearance    `json:"appearance"`
}

// CharacterItems represents World of Warcraft character items.
type CharacterItems struct {
	AverageItemLevel         int                 `json:"averageItemLevel"`
	AverageItemLevelEquipped int                 `json:"averageItemLevelEquipped"`
	Head                     CharacterItemSlot   `json:"head"`
	Neck                     CharacterItemSlot   `json:"neck"`
	Shoulder                 CharacterItemSlot   `json:"shoulder"`
	Back                     CharacterItemSlot   `json:"back"`
	Chest                    CharacterItemSlot   `json:"chest"`
	Shirt                    CharacterItemSlot   `json:"shirt"`
	Wrist                    CharacterItemSlot   `json:"wrist"`
	Hands                    CharacterItemSlot   `json:"hands"`
	Waist                    CharacterItemSlot   `json:"waist"`
	Legs                     CharacterItemSlot   `json:"legs"`
	Feet                     CharacterItemSlot   `json:"feet"`
	Finger1                  CharacterItemSlot   `json:"finger1"`
	Finger2                  CharacterItemSlot   `json:"finger2"`
	Trinket1                 CharacterItemSlot   `json:"trinket1"`
	Trinket2                 CharacterItemSlot   `json:"trinket2"`
	MainHand                 CharacterWeaponSlot `json:"mainHand"`
}

// CharacterStats represents World of Warcraft character stats.
type CharacterStats struct {
	Health                      int     `json:"health"`
	PowerType                   string  `json:"powerType"`
	Power                       int     `json:"power"`
	Str                         int     `json:"str"`
	Agi                         int     `json:"agi"`
	Int                         int     `json:"int"`
	Sta                         int     `json:"sta"`
	SpeedRating                 float64 `json:"speedRating"`
	SpeedRatingBonus            int     `json:"speedRatingBonus"`
	Crit                        float64 `json:"crit"`
	CritRating                  int     `json:"critRating"`
	Haste                       float64 `json:"haste"`
	HasteRating                 int     `json:"hasteRating"`
	HasteRatingPercent          float64 `json:"hasteRatingPercent"`
	Mastery                     float64 `json:"mastery"`
	MasteryRating               int     `json:"masteryRating"`
	Leech                       int     `json:"leech"`
	LeechRating                 int     `json:"leechRating"`
	LeechRatingBonus            int     `json:"leechRatingBonus"`
	Versatility                 int     `json:"versatility"`
	VersatilityDamageDoneBonus  float64 `json:"versatilityDamageDoneBonus"`
	VersatilityHealingDoneBonus float64 `json:"versatilityHealingDoneBonus"`
	VersatilityDamageTakenBonus float64 `json:"versatilityDamageTakenBonus"`
	AvoidanceRating             int     `json:"avoidanceRating"`
	AvoidanceRatingBonus        int     `json:"avoidanceRatingBonus"`
	SpellPen                    int     `json:"spellPen"`
	SpellCrit                   float64 `json:"spellCrit"`
	SpellCritRating             int     `json:"spellCritRating"`
	Mana5                       int     `json:"mana5"`
	Mana5Combat                 int     `json:"mana5Combat"`
	Armor                       int     `json:"armor"`
	Dodge                       float64 `json:"dodge"`
	DodgeRating                 int     `json:"dodgeRating"`
	Parry                       int     `json:"parry"`
	ParryRating                 int     `json:"parryRating"`
	Block                       int     `json:"block"`
	BlockRating                 int     `json:"blockRating"`
	MainHandDmgMin              int     `json:"mainHandDmgMin"`
	MainHandDmgMax              int     `json:"mainHandDmgMax"`
	MainHandSpeed               float64 `json:"mainHandSpeed"`
	MainHandDps                 float64 `json:"mainHandDps"`
	OffHandDmgMin               int     `json:"offHandDmgMin"`
	OffHandDmgMax               int     `json:"offHandDmgMax"`
	OffHandSpeed                float64 `json:"offHandSpeed"`
	OffHandDps                  float64 `json:"offHandDps"`
	RangedDmgMin                int     `json:"rangedDmgMin"`
	RangedDmgMax                int     `json:"rangedDmgMax"`
	RangedSpeed                 int     `json:"rangedSpeed"`
	RangedDps                   int     `json:"rangedDps"`
}

// CharacterProfession represents a World of Warcraft character profession.
type CharacterProfession struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Icon    string        `json:"icon"`
	Rank    int           `json:"rank"`
	Max     int           `json:"max"`
	Recipes []interface{} `json:"recipes"`
}

// CharacterProfessions represents World of Warcraft character professions.
type CharacterProfessions struct {
	Primary   []CharacterProfession `json:"primary"`
	Secondary []CharacterProfession `json:"secondary"`
}

// CharacterReputation represents a World of Warcraft character reputation.
type CharacterReputation struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Standing int    `json:"standing"`
	Value    int    `json:"value"`
	Max      int    `json:"max"`
}

// CharacterTitle represents a World of Warcraft character title.
type CharacterTitle struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Selected bool   `json:"selected,omitempty"`
}

// CharacterAchievements represents World of Warcraft character achievements.
type CharacterAchievements struct {
	AchievementsCompleted          []int         `json:"achievementsCompleted"`
	AchievementsCompletedTimestamp []interface{} `json:"achievementsCompletedTimestamp"`
	Criteria                       []int         `json:"criteria"`
	CriteriaQuantity               []interface{} `json:"criteriaQuantity"`
	CriteriaTimestamp              []int64       `json:"criteriaTimestamp"`
	CriteriaCreated                []interface{} `json:"criteriaCreated"`
}

// CharacterStatistic represents a World of Warcraft character statistic.
type CharacterStatistic struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	LastUpdated int    `json:"lastUpdated"`
	Money       bool   `json:"money"`
	Highest     string `json:"highest,omitempty"`
}

// CharacterInnerSubcategory represents a World of Warcraft character inner subcategory.
type CharacterInnerSubcategory struct {
	ID         int                  `json:"id"`
	Name       string               `json:"name"`
	Statistics []CharacterStatistic `json:"statistics"`
}

// CharacterSubcategory represents a World of Warcraft character subcategory.
type CharacterSubcategory struct {
	ID            int                         `json:"id"`
	Name          string                      `json:"name"`
	Statistics    []CharacterStatistic        `json:"statistics"`
	SubCategories []CharacterInnerSubcategory `json:"subCategories,omitempty"`
}

// CharacterStatistics represents World of Warcraft character statistics.
type CharacterStatistics struct {
	ID            int                    `json:"id"`
	Name          string                 `json:"name"`
	SubCategories []CharacterSubcategory `json:"subCategories"`
}

// CharacterSpell represents a World of Warcraft character spell.
type CharacterSpell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Range       string `json:"range"`
	PowerCost   string `json:"powerCost"`
	CastTime    string `json:"castTime"`
	Cooldown    string `json:"cooldown"`
}

// CharacterTalent represents a World of Warcraft character talent.
type CharacterTalent struct {
	Tier   int            `json:"tier"`
	Column int            `json:"column"`
	Spell  CharacterSpell `json:"spell"`
	Spec   CharacterSpec  `json:"spec,omitempty"`
}

// CharacterTalents represents World of Warcraft character talents.
type CharacterTalents struct {
	Selected   bool              `json:"selected,omitempty"`
	Talents    []CharacterTalent `json:"talents"`
	Spec       CharacterSpec     `json:"spec,omitempty"`
	CalcTalent string            `json:"calcTalent"`
	CalcSpec   string            `json:"calcSpec"`
}

// CharacterAppearance represents a World of Warcraft character appearance.
type CharacterAppearance struct {
	FaceVariation        int   `json:"faceVariation"`
	SkinColor            int   `json:"skinColor"`
	HairVariation        int   `json:"hairVariation"`
	HairColor            int   `json:"hairColor"`
	FeatureVariation     int   `json:"featureVariation"`
	ShowHelm             bool  `json:"showHelm"`
	ShowCloak            bool  `json:"showCloak"`
	CustomDisplayOptions []int `json:"customDisplayOptions"`
}

// CharacterMount represents a World of Warcraft character mount.
type CharacterMount struct {
	Name       string `json:"name"`
	SpellID    int    `json:"spellId"`
	CreatureID int    `json:"creatureId"`
	ItemID     int    `json:"itemId"`
	QualityID  int    `json:"qualityId"`
	Icon       string `json:"icon"`
	IsGround   bool   `json:"isGround"`
	IsFlying   bool   `json:"isFlying"`
	IsAquatic  bool   `json:"isAquatic"`
	IsJumping  bool   `json:"isJumping"`
}

// CharacterMounts represents World of Warcraft character mounts.
type CharacterMounts struct {
	NumCollected    int              `json:"numCollected"`
	NumNotCollected int              `json:"numNotCollected"`
	Collected       []CharacterMount `json:"collected"`
}

// CharacterPetStat represents a World of Warcraft character pet stat.
type CharacterPetStat struct {
	SpeciesID    int `json:"speciesId"`
	BreedID      int `json:"breedId"`
	PetQualityID int `json:"petQualityId"`
	Level        int `json:"level"`
	Health       int `json:"health"`
	Power        int `json:"power"`
	Speed        int `json:"speed"`
}

// CharacterPet represents a World of Warcraft character pet.
type CharacterPet struct {
	Name                        string           `json:"name"`
	SpellID                     int              `json:"spellId"`
	CreatureID                  int              `json:"creatureId"`
	ItemID                      int              `json:"itemId"`
	QualityID                   int              `json:"qualityId"`
	Icon                        string           `json:"icon"`
	Stats                       CharacterPetStat `json:"stats"`
	BattlePetGUID               string           `json:"battlePetGuid"`
	IsFavorite                  bool             `json:"isFavorite"`
	IsFirstAbilitySlotSelected  bool             `json:"isFirstAbilitySlotSelected"`
	IsSecondAbilitySlotSelected bool             `json:"isSecondAbilitySlotSelected"`
	IsThirdAbilitySlotSelected  bool             `json:"isThirdAbilitySlotSelected"`
	CreatureName                string           `json:"creatureName"`
	CanBattle                   bool             `json:"canBattle"`
}

// CharacterPets represents World of Warcraft character pets.
type CharacterPets struct {
	NumCollected    int            `json:"numCollected"`
	NumNotCollected int            `json:"numNotCollected"`
	Collected       []CharacterPet `json:"collected"`
}

// CharacterPetSlot represents a World of Warcraft character pet slot.
type CharacterPetSlot struct {
	Slot      int           `json:"slot"`
	IsEmpty   bool          `json:"isEmpty"`
	IsLocked  bool          `json:"isLocked"`
	Abilities []interface{} `json:"abilities"`
}

// CharacterRaidBoss represents a World of Warcraft character raidboss.
type CharacterRaidBoss struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	NormalKills     int    `json:"normalKills"`
	NormalTimestamp int    `json:"normalTimestamp"`
}

// CharacterRaid represents a World of Warcraft character raid.
type CharacterRaid struct {
	Name   string              `json:"name"`
	Lfr    int                 `json:"lfr"`
	Normal int                 `json:"normal"`
	Heroic int                 `json:"heroic"`
	Mythic int                 `json:"mythic"`
	ID     int                 `json:"id"`
	Bosses []CharacterRaidBoss `json:"bosses"`
}

// CharacterProgression represents a World of Warcraft character's progression.
type CharacterProgression struct {
	Raids []CharacterRaid `json:"raids"`
}

// CharacterArenaBracket represents a World of Warcraft character arena bracket.
type CharacterArenaBracket struct {
	Slug         string `json:"slug"`
	Rating       int    `json:"rating"`
	WeeklyPlayed int    `json:"weeklyPlayed"`
	WeeklyWon    int    `json:"weeklyWon"`
	WeeklyLost   int    `json:"weeklyLost"`
	SeasonPlayed int    `json:"seasonPlayed"`
	SeasonWon    int    `json:"seasonWon"`
	SeasonLost   int    `json:"seasonLost"`
}

// CharacterArenaBrackets represents World of Warcraft character arena brackets.
type CharacterArenaBrackets struct {
	ARENABRACKET2V2         CharacterArenaBracket `json:"ARENA_BRACKET_2v2"`
	ARENABRACKET3V3         CharacterArenaBracket `json:"ARENA_BRACKET_3v3"`
	ARENABRACKETRBG         CharacterArenaBracket `json:"ARENA_BRACKET_RBG"`
	ARENABRACKET2V2SKIRMISH CharacterArenaBracket `json:"ARENA_BRACKET_2v2_SKIRMISH"`
	UNKNOWN                 CharacterArenaBracket `json:"UNKNOWN"`
}

// CharacterPvp represents a World of Warcraft character PvP.
type CharacterPvp struct {
	Brackets CharacterArenaBrackets `json:"brackets"`
}

// CharacterAuditSlot represents a World of Warcraft character audit slots.
type CharacterAuditSlot struct {
	Num2  int `json:"2"`
	Num4  int `json:"4"`
	Num5  int `json:"5"`
	Num6  int `json:"6"`
	Num7  int `json:"7"`
	Num8  int `json:"8"`
	Num9  int `json:"9"`
	Num14 int `json:"14"`
	Num15 int `json:"15"`
}

// CharacterAuditUnenchantedItems represents unenchanted items from a character audit.
type CharacterAuditUnenchantedItems struct {
	Num2  int `json:"2"`
	Num4  int `json:"4"`
	Num6  int `json:"6"`
	Num7  int `json:"7"`
	Num8  int `json:"8"`
	Num9  int `json:"9"`
	Num14 int `json:"14"`
	Num15 int `json:"15"`
}

// CharacterAuditItemsWithEmptySockets represents items with empty sockets from a character audit.
type CharacterAuditItemsWithEmptySockets struct {
}

// CharacterAuditInappropriateArmorType represents inappropriate armor types from a character audit.
type CharacterAuditInappropriateArmorType struct {
}

// CharacterAuditLowLevelItems represents low level items returned from a character audit.
type CharacterAuditLowLevelItems struct {
}

// CharacterAuditMissingExtraSockets represents missing extra sockets from a character audit.
type CharacterAuditMissingExtraSockets struct {
	Num5 int `json:"5"`
}

// CharacterAuditMissingBlacksmithSockets represents missing blacksmith sockets from a character audit.
type CharacterAuditMissingBlacksmithSockets struct {
}

// CharacterAuditMissingEnchanterEnchants represents missing enchanter enchants from a character audit.
type CharacterAuditMissingEnchanterEnchants struct {
}

// CharacterAuditMissingEngineerEnchants represents missing engineer enchants from a character audit.
type CharacterAuditMissingEngineerEnchants struct {
}

// CharacterAuditMissingScribeEnchants represents missing scribe enchants from a character audit.
type CharacterAuditMissingScribeEnchants struct {
}

// CharacterAuditMissingLeatherworkerEnchants represents missing leatherworker enchants.
type CharacterAuditMissingLeatherworkerEnchants struct {
}

// CharacterAuditSpell represents a World of Warcraft character audit spell.
type CharacterAuditSpell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
}

// CharacterAuditItemSpell represents a World of Warcraft character audit item spell.
type CharacterAuditItemSpell struct {
	SpellID           int                 `json:"spellId"`
	Spell             CharacterAuditSpell `json:"spell"`
	NCharges          int                 `json:"nCharges"`
	Consumable        bool                `json:"consumable"`
	CategoryID        int                 `json:"categoryId"`
	Trigger           string              `json:"trigger"`
	ScaledDescription string              `json:"scaledDescription"`
}

// CharacterAuditItemSource represents a World of Warcraft character audit item source.
type CharacterAuditItemSource struct {
	SourceID   int    `json:"sourceId"`
	SourceType string `json:"sourceType"`
}

// CharacterAuditBonusSummary represents a World of Warcraft character audit bonus summary.
type CharacterAuditBonusSummary struct {
	DefaultBonusLists []interface{} `json:"defaultBonusLists"`
	ChanceBonusLists  []interface{} `json:"chanceBonusLists"`
	BonusChances      []interface{} `json:"bonusChances"`
}

// CharacterAuditRecommendedBeltBuckle represents the recommended belt buckle.
type CharacterAuditRecommendedBeltBuckle struct {
	ID                   int                        `json:"id"`
	Description          string                     `json:"description"`
	Name                 string                     `json:"name"`
	Icon                 string                     `json:"icon"`
	Stackable            int                        `json:"stackable"`
	ItemBind             int                        `json:"itemBind"`
	BonusStats           []interface{}              `json:"bonusStats"`
	ItemSpells           []CharacterAuditItemSpell  `json:"itemSpells"`
	BuyPrice             int                        `json:"buyPrice"`
	ItemClass            int                        `json:"itemClass"`
	ItemSubClass         int                        `json:"itemSubClass"`
	ContainerSlots       int                        `json:"containerSlots"`
	InventoryType        int                        `json:"inventoryType"`
	Equippable           bool                       `json:"equippable"`
	ItemLevel            int                        `json:"itemLevel"`
	MaxCount             int                        `json:"maxCount"`
	MaxDurability        int                        `json:"maxDurability"`
	MinFactionID         int                        `json:"minFactionId"`
	MinReputation        int                        `json:"minReputation"`
	Quality              int                        `json:"quality"`
	SellPrice            int                        `json:"sellPrice"`
	RequiredSkill        int                        `json:"requiredSkill"`
	RequiredLevel        int                        `json:"requiredLevel"`
	RequiredSkillRank    int                        `json:"requiredSkillRank"`
	ItemSource           CharacterAuditItemSource   `json:"itemSource"`
	BaseArmor            int                        `json:"baseArmor"`
	HasSockets           bool                       `json:"hasSockets"`
	IsAuctionable        bool                       `json:"isAuctionable"`
	Armor                int                        `json:"armor"`
	DisplayInfoID        int                        `json:"displayInfoId"`
	NameDescription      string                     `json:"nameDescription"`
	NameDescriptionColor string                     `json:"nameDescriptionColor"`
	Upgradable           bool                       `json:"upgradable"`
	HeroicTooltip        bool                       `json:"heroicTooltip"`
	Context              string                     `json:"context"`
	BonusLists           []interface{}              `json:"bonusLists"`
	AvailableContexts    []string                   `json:"availableContexts"`
	BonusSummary         CharacterAuditBonusSummary `json:"bonusSummary"`
	ArtifactID           int                        `json:"artifactId"`
}

// CharacterAudit represents a World of Warcraft character audit.
type CharacterAudit struct {
	NumberOfIssues               int                                        `json:"numberOfIssues"`
	Slots                        CharacterAuditSlot                         `json:"slots"`
	EmptyGlyphSlots              int                                        `json:"emptyGlyphSlots"`
	UnspentTalentPoints          int                                        `json:"unspentTalentPoints"`
	NoSpec                       bool                                       `json:"noSpec"`
	UnenchantedItems             CharacterAuditUnenchantedItems             `json:"unenchantedItems"`
	EmptySockets                 int                                        `json:"emptySockets"`
	ItemsWithEmptySockets        CharacterAuditItemsWithEmptySockets        `json:"itemsWithEmptySockets"`
	AppropriateArmorType         int                                        `json:"appropriateArmorType"`
	InappropriateArmorType       CharacterAuditInappropriateArmorType       `json:"inappropriateArmorType"`
	LowLevelItems                CharacterAuditLowLevelItems                `json:"lowLevelItems"`
	LowLevelThreshold            int                                        `json:"lowLevelThreshold"`
	MissingExtraSockets          CharacterAuditMissingExtraSockets          `json:"missingExtraSockets"`
	RecommendedBeltBuckle        CharacterAuditRecommendedBeltBuckle        `json:"recommendedBeltBuckle"`
	MissingBlacksmithSockets     CharacterAuditMissingBlacksmithSockets     `json:"missingBlacksmithSockets"`
	MissingEnchanterEnchants     CharacterAuditMissingEnchanterEnchants     `json:"missingEnchanterEnchants"`
	MissingEngineerEnchants      CharacterAuditMissingEngineerEnchants      `json:"missingEngineerEnchants"`
	MissingScribeEnchants        CharacterAuditMissingScribeEnchants        `json:"missingScribeEnchants"`
	NMissingJewelcrafterGems     int                                        `json:"nMissingJewelcrafterGems"`
	MissingLeatherworkerEnchants CharacterAuditMissingLeatherworkerEnchants `json:"missingLeatherworkerEnchants"`
}

// CharacterData represents World of Warcraft character data.
type CharacterData struct {
	LastModified        int64                 `json:"lastModified"`
	Name                string                `json:"name"`
	Realm               string                `json:"realm"`
	Battlegroup         string                `json:"battlegroup"`
	Class               int                   `json:"class"`
	Race                int                   `json:"race"`
	Gender              int                   `json:"gender"`
	Level               int                   `json:"level"`
	AchievementPoints   int                   `json:"achievementPoints"`
	Thumbnail           string                `json:"thumbnail"`
	CalcClass           string                `json:"calcClass"`
	Faction             int                   `json:"faction"`
	Guild               CharacterGuild        `json:"guild"`
	Feed                []CharacterFeedItem   `json:"feed"`
	Items               CharacterItems        `json:"items"`
	Stats               CharacterStats        `json:"stats"`
	Professions         CharacterProfessions  `json:"professions"`
	Reputation          []CharacterReputation `json:"reputation"`
	Titles              []CharacterTitle      `json:"titles"`
	Achievements        CharacterAchievements `json:"achievements"`
	Statistics          CharacterStatistics   `json:"statistics"`
	Talents             []CharacterTalent     `json:"talents"`
	Appearance          CharacterAppearance   `json:"appearance"`
	Mounts              CharacterMounts       `json:"mounts"`
	Pets                CharacterPets         `json:"pets"`
	PetSlots            []CharacterPetSlot    `json:"petSlots"`
	Progression         CharacterProgression  `json:"progression"`
	Pvp                 CharacterPvp          `json:"pvp"`
	Quests              []int                 `json:"quests"`
	Audit               CharacterAudit        `json:"audit"`
	TotalHonorableKills int                   `json:"totalHonorableKills"`
}
