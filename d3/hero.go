package d3

// Dye represents a Diablo III dye.
type SkillSet struct {
	Skill Skill `json:"skill"`
	Rune  Skill `json:"rune"`
}

// Dye represents a Diablo III dye.
type HeroSkills struct {
	Active  []SkillSet `json:"active"`
	Passive []SkillSet `json:"passive"`
}

// Dye represents a Diablo III dye.
type DyeColor struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Icon          string `json:"icon"`
	TooltipParams string `json:"tooltipParams"`
}

// Dye represents a Diablo III dye.
type TransmogItem struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Icon          string `json:"icon"`
	DisplayColor  string `json:"displayColor"`
	TooltipParams string `json:"tooltipParams"`
}

// Dye represents a Diablo III dye.
type HeroItemData struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Icon          string       `json:"icon"`
	DisplayColor  string       `json:"displayColor"`
	TooltipParams string       `json:"tooltipParams"`
	DyeColor      DyeColor     `json:"dyeColor"`
	TransmogItem  TransmogItem `json:"transmogItem"`
}

// Dye represents a Diablo III dye.
type HeroItems struct {
	Head        HeroItemData `json:"head"`
	Neck        HeroItemData `json:"neck"`
	Torso       HeroItemData `json:"torso"`
	Shoulders   HeroItemData `json:"shoulders"`
	Legs        HeroItemData `json:"legs"`
	Waist       HeroItemData `json:"waist"`
	Hands       HeroItemData `json:"hands"`
	Bracers     HeroItemData `json:"bracers"`
	Feet        HeroItemData `json:"feet"`
	LeftFinger  HeroItemData `json:"leftFinger"`
	RightFinger HeroItemData `json:"rightFinger"`
	MainHand    HeroItemData `json:"mainHand"`
	OffHand     HeroItemData `json:"offHand"`
}

// FollowerHands represents a Diablo III follwer's hands.
type FollowerHands struct {
	MainHand TransmogItem `json:"mainHand"`
	OffHand  TransmogItem `json:"offHand"`
}

// FollowerStats represents Diablo III follower stats.
type FollowerStats struct {
	GoldFind        int `json:"goldFind"`
	MagicFind       int `json:"magicFind"`
	ExperienceBonus int `json:"experienceBonus"`
}

// HeroFollower represents a Diablo III hero follower.
type HeroFollower struct {
	Slug   string        `json:"slug"`
	Level  int           `json:"level"`
	Items  FollowerHands `json:"items"`
	Stats  FollowerStats `json:"stats"`
	Skills []interface{} `json:"skills"`
}

// HeroFollowers represents Diablo III hero followers.
type HeroFollowers struct {
	Templar     HeroFollower `json:"templar"`
	Scoundrel   HeroFollower `json:"scoundrel"`
	Enchantress HeroFollower `json:"enchantress"`
}

// HeroCompletedQuest represents a completed Diablo III hero quest.
type HeroCompletedQuest struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// HeroAct represents a Diablo III hero act.
type HeroAct struct {
	Completed       bool                 `json:"completed"`
	CompletedQuests []HeroCompletedQuest `json:"completedQuests"`
}

// HeroProgression represents a Diablo III hero's progression.
type HeroProgression struct {
	Act1 HeroAct `json:"act1"`
	Act2 HeroAct `json:"act2"`
	Act3 HeroAct `json:"act3"`
	Act4 HeroAct `json:"act4"`
	Act5 HeroAct `json:"act5"`
}

// HeroStats represents Diablo III hero stats.
type HeroStats struct {
	Life              float64 `json:"life"`
	Damage            float64 `json:"damage"`
	Toughness         int     `json:"toughness"`
	Healing           int     `json:"healing"`
	AttackSpeed       float64 `json:"attackSpeed"`
	Armor             int     `json:"armor"`
	Strength          int     `json:"strength"`
	Dexterity         int     `json:"dexterity"`
	Vitality          int     `json:"vitality"`
	Intelligence      int     `json:"intelligence"`
	PhysicalResist    int     `json:"physicalResist"`
	FireResist        int     `json:"fireResist"`
	ColdResist        int     `json:"coldResist"`
	LightningResist   int     `json:"lightningResist"`
	PoisonResist      int     `json:"poisonResist"`
	ArcaneResist      int     `json:"arcaneResist"`
	BlockChance       int     `json:"blockChance"`
	BlockAmountMin    int     `json:"blockAmountMin"`
	BlockAmountMax    int     `json:"blockAmountMax"`
	GoldFind          float64 `json:"goldFind"`
	CritChance        float64 `json:"critChance"`
	Thorns            int     `json:"thorns"`
	LifeSteal         int     `json:"lifeSteal"`
	LifePerKill       int     `json:"lifePerKill"`
	LifeOnHit         int     `json:"lifeOnHit"`
	PrimaryResource   int     `json:"primaryResource"`
	SecondaryResource int     `json:"secondaryResource"`
}

// Hero represents a Diablo III hero.
type Hero struct {
	ID                       int             `json:"id"`
	Name                     string          `json:"name"`
	Class                    string          `json:"class"`
	Gender                   int             `json:"gender"`
	Level                    int             `json:"level"`
	ParagonLevel             int             `json:"paragonLevel"`
	Kills                    Kills           `json:"kills"`
	Hardcore                 bool            `json:"hardcore"`
	Seasonal                 bool            `json:"seasonal"`
	SeasonCreated            int             `json:"seasonCreated"`
	Skills                   HeroSkills      `json:"skills"`
	Items                    HeroItems       `json:"items"`
	Followers                HeroFollowers   `json:"followers"`
	LegendaryPowers          []TransmogItem  `json:"legendaryPowers"`
	Progression              HeroProgression `json:"progression"`
	Alive                    bool            `json:"alive"`
	LastUpdated              int             `json:"lastUpdated"`
	HighestSoloRiftCompleted int             `json:"highestSoloRiftCompleted"`
	Stats                    HeroStats       `json:"stats"`
}
