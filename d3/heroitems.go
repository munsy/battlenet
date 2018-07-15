package d3

// Get API Detailed Hero Items
type DetailedHeroItems struct {
	Head struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            int    `json:"armor"`
		AttacksPerSecond int    `json:"attacksPerSecond"`
		MinDamage        int    `json:"minDamage"`
		MaxDamage        int    `json:"maxDamage"`
		Slots            string `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			Attributes []string `json:"attributes"`
			IsGem      bool     `json:"isGem"`
			IsJewel    bool     `json:"isJewel"`
		} `json:"gems"`
		SeasonRequiredToDrop int `json:"seasonRequiredToDrop"`
		Dye                  struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"dye"`
		Transmog               HeroItemData `json:"transmog"`
		IsSeasonRequiredToDrop bool         `json:"isSeasonRequiredToDrop"`
	} `json:"head"`
	Neck struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            int    `json:"armor"`
		AttacksPerSecond int    `json:"attacksPerSecond"`
		MinDamage        int    `json:"minDamage"`
		MaxDamage        int    `json:"maxDamage"`
		Slots            string `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			JewelRank                int      `json:"jewelRank"`
			JewelSecondaryUnlockRank int      `json:"jewelSecondaryUnlockRank"`
			Attributes               []string `json:"attributes"`
			IsGem                    bool     `json:"isGem"`
			IsJewel                  bool     `json:"isJewel"`
		} `json:"gems"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"neck"`
	Torso struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            int    `json:"armor"`
		AttacksPerSecond int    `json:"attacksPerSecond"`
		MinDamage        int    `json:"minDamage"`
		MaxDamage        int    `json:"maxDamage"`
		Slots            string `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			Attributes []string `json:"attributes"`
			IsGem      bool     `json:"isGem"`
			IsJewel    bool     `json:"isJewel"`
		} `json:"gems"`
		Set struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop int `json:"seasonRequiredToDrop"`
		Dye                  struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"dye"`
		Transmog               HeroItemData `json:"transmog"`
		IsSeasonRequiredToDrop bool         `json:"isSeasonRequiredToDrop"`
	} `json:"torso"`
	Shoulders struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond int     `json:"attacksPerSecond"`
		MinDamage        int     `json:"minDamage"`
		MaxDamage        int     `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets          int `json:"openSockets"`
		SeasonRequiredToDrop int `json:"seasonRequiredToDrop"`
		Dye                  struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"dye"`
		Transmog               HeroItemData `json:"transmog"`
		IsSeasonRequiredToDrop bool         `json:"isSeasonRequiredToDrop"`
	} `json:"shoulders"`
	Legs struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            int    `json:"armor"`
		AttacksPerSecond int    `json:"attacksPerSecond"`
		MinDamage        int    `json:"minDamage"`
		MaxDamage        int    `json:"maxDamage"`
		Slots            string `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			Attributes []string `json:"attributes"`
			IsGem      bool     `json:"isGem"`
			IsJewel    bool     `json:"isJewel"`
		} `json:"gems"`
		Set struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop int `json:"seasonRequiredToDrop"`
		Dye                  struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"dye"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"legs"`
	Waist struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            int    `json:"armor"`
		AttacksPerSecond int    `json:"attacksPerSecond"`
		MinDamage        int    `json:"minDamage"`
		MaxDamage        int    `json:"maxDamage"`
		Slots            string `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Set         struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"waist"`
	Hands struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond int     `json:"attacksPerSecond"`
		MinDamage        int     `json:"minDamage"`
		MaxDamage        int     `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Set         struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop int `json:"seasonRequiredToDrop"`
		Dye                  struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"dye"`
		Transmog               HeroItemData `json:"transmog"`
		IsSeasonRequiredToDrop bool         `json:"isSeasonRequiredToDrop"`
	} `json:"hands"`
	Bracers struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond int     `json:"attacksPerSecond"`
		MinDamage        int     `json:"minDamage"`
		MaxDamage        int     `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets            int  `json:"openSockets"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"bracers"`
	Feet struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond int     `json:"attacksPerSecond"`
		MinDamage        int     `json:"minDamage"`
		MaxDamage        int     `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets          int `json:"openSockets"`
		SeasonRequiredToDrop int `json:"seasonRequiredToDrop"`
		Dye                  struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"dye"`
		Transmog               HeroItemData `json:"transmog"`
		IsSeasonRequiredToDrop bool         `json:"isSeasonRequiredToDrop"`
	} `json:"feet"`
	LeftFinger struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            int    `json:"armor"`
		AttacksPerSecond int    `json:"attacksPerSecond"`
		MinDamage        int    `json:"minDamage"`
		MaxDamage        int    `json:"maxDamage"`
		Slots            string `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			JewelRank                int      `json:"jewelRank"`
			JewelSecondaryUnlockRank int      `json:"jewelSecondaryUnlockRank"`
			Attributes               []string `json:"attributes"`
			IsGem                    bool     `json:"isGem"`
			IsJewel                  bool     `json:"isJewel"`
		} `json:"gems"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"leftFinger"`
	RightFinger struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            int    `json:"armor"`
		AttacksPerSecond int    `json:"attacksPerSecond"`
		MinDamage        int    `json:"minDamage"`
		MaxDamage        int    `json:"maxDamage"`
		Slots            string `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets            int  `json:"openSockets"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"rightFinger"`
	MainHand struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            int     `json:"armor"`
		Damage           string  `json:"damage"`
		Dps              string  `json:"dps"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        int     `json:"minDamage"`
		MaxDamage        int     `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			Attributes []string `json:"attributes"`
			IsGem      bool     `json:"isGem"`
			IsJewel    bool     `json:"isJewel"`
		} `json:"gems"`
		Set struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"mainHand"`
	OffHand struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            int     `json:"armor"`
		Damage           string  `json:"damage"`
		Dps              string  `json:"dps"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        int     `json:"minDamage"`
		MaxDamage        int     `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			Attributes []string `json:"attributes"`
			IsGem      bool     `json:"isGem"`
			IsJewel    bool     `json:"isJewel"`
		} `json:"gems"`
		Set struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"offHand"`
}
