package wow

type Battlegroup struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type BattleGroupsData struct {
	Battlegroups []Battlegroup `json:"battlegroups"`
}

type CharacterRace struct {
	ID   int    `json:"id"`
	Mask int    `json:"mask"`
	Side string `json:"side"`
	Name string `json:"name"`
}

type CharacterRacesData struct {
	Races []CharacterRace `json:"races"`
}

type CharacterClass struct {
	ID        int    `json:"id"`
	Mask      int    `json:"mask"`
	PowerType string `json:"powerType"`
	Name      string `json:"name"`
}

type CharacterClassesData struct {
	Classes []CharacterClass `json:"classes"`
}

type CharacterAchievementCriteria struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	OrderIndex  int    `json:"orderIndex"`
	Max         int    `json:"max"`
}

type CharacterAchievement struct {
	ID          int                            `json:"id"`
	Title       string                         `json:"title"`
	Points      int                            `json:"points"`
	Description string                         `json:"description"`
	RewardItems []interface{}                  `json:"rewardItems"`
	Icon        string                         `json:"icon"`
	Criteria    []CharacterAchievementCriteria `json:"criteria"`
	AccountWide bool                           `json:"accountWide"`
	FactionID   int                            `json:"factionId"`
}

type CharacterAchievementCategory struct {
	ID           int                    `json:"id"`
	Achievements []CharacterAchievement `json:"achievements"`
	Name         string                 `json:"name"`
}

type CharacterAchievements struct {
	ID           int                            `json:"id"`
	Achievements []CharacterAchievement         `json:"achievements,omitempty"`
	Name         string                         `json:"name"`
	Categories   []CharacterAchievementCategory `json:"categories,omitempty"`
}

type CharacterAchievementData struct {
	Achievements []CharacterAchievements `json:"achievements"`
}

type GuildRewardsData struct {
	Rewards []struct {
		MinGuildLevel    int `json:"minGuildLevel"`
		MinGuildRepLevel int `json:"minGuildRepLevel"`
		Achievement      struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Points      int    `json:"points"`
			Description string `json:"description"`
			Reward      string `json:"reward"`
			RewardItems []struct {
				ID            int    `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				Quality       int    `json:"quality"`
				ItemLevel     int    `json:"itemLevel"`
				TooltipParams struct {
					TimewalkerLevel   int `json:"timewalkerLevel"`
					AzeritePower0     int `json:"azeritePower0"`
					AzeritePower1     int `json:"azeritePower1"`
					AzeritePower2     int `json:"azeritePower2"`
					AzeritePower3     int `json:"azeritePower3"`
					AzeritePowerLevel int `json:"azeritePowerLevel"`
				} `json:"tooltipParams"`
				Stats []struct {
					Stat   int `json:"stat"`
					Amount int `json:"amount"`
				} `json:"stats"`
				Armor                int           `json:"armor"`
				Context              string        `json:"context"`
				BonusLists           []interface{} `json:"bonusLists"`
				ArtifactID           int           `json:"artifactId"`
				DisplayInfoID        int           `json:"displayInfoId"`
				ArtifactAppearanceID int           `json:"artifactAppearanceId"`
				ArtifactTraits       []interface{} `json:"artifactTraits"`
				Relics               []interface{} `json:"relics"`
				Appearance           struct {
				} `json:"appearance"`
			} `json:"rewardItems"`
			Icon     string `json:"icon"`
			Criteria []struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				OrderIndex  int    `json:"orderIndex"`
				Max         int    `json:"max"`
			} `json:"criteria"`
			AccountWide bool `json:"accountWide"`
			FactionID   int  `json:"factionId"`
		} `json:"achievement,omitempty"`
		Item struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []interface{} `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
		} `json:"item"`
		Races []int `json:"races,omitempty"`
	} `json:"rewards"`
}

type GuildPerksData struct {
	Perks []struct {
		GuildLevel int `json:"guildLevel"`
		Spell      struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Subtext     string `json:"subtext"`
			Icon        string `json:"icon"`
			Description string `json:"description"`
			CastTime    string `json:"castTime"`
		} `json:"spell"`
	} `json:"perks"`
}

type GuildAchievementsData struct {
	Achievements []struct {
		ID           int `json:"id"`
		Achievements []struct {
			ID          int           `json:"id"`
			Title       string        `json:"title"`
			Points      int           `json:"points"`
			Description string        `json:"description"`
			RewardItems []interface{} `json:"rewardItems"`
			Icon        string        `json:"icon"`
			Criteria    []struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				OrderIndex  int    `json:"orderIndex"`
				Max         int    `json:"max"`
			} `json:"criteria"`
			AccountWide bool   `json:"accountWide"`
			FactionID   int    `json:"factionId"`
			Reward      string `json:"reward,omitempty"`
		} `json:"achievements"`
		Name       string `json:"name"`
		Categories []struct {
			ID           int `json:"id"`
			Achievements []struct {
				ID          int           `json:"id"`
				Title       string        `json:"title"`
				Points      int           `json:"points"`
				Description string        `json:"description"`
				RewardItems []interface{} `json:"rewardItems"`
				Icon        string        `json:"icon"`
				Criteria    []struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					OrderIndex  int    `json:"orderIndex"`
					Max         int    `json:"max"`
				} `json:"criteria"`
				AccountWide bool `json:"accountWide"`
				FactionID   int  `json:"factionId"`
			} `json:"achievements"`
			Name string `json:"name"`
		} `json:"categories,omitempty"`
	} `json:"achievements"`
}

type ItemClassesData struct {
	Classes []struct {
		Class      int    `json:"class"`
		Name       string `json:"name"`
		Subclasses []struct {
			Subclass int    `json:"subclass"`
			Name     string `json:"name"`
		} `json:"subclasses"`
	} `json:"classes"`
}

type TalentsData struct {
	Num1 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"1"`
	Num2 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"2"`
	Num3 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		PetSpecs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"petSpecs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"3"`
	Num4 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"4"`
	Num5 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec,omitempty"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"5"`
	Num6 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"6"`
	Num7 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"7"`
	Num8 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"8"`
	Num9 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"9"`
	Num10 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec,omitempty"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"10"`
	Num11 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"11"`
	Num12 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"12"`
}

type PetTypes struct {
	ID              int    `json:"id"`
	Key             string `json:"key"`
	Name            string `json:"name"`
	TypeAbilityID   int    `json:"typeAbilityId"`
	StrongAgainstID int    `json:"strongAgainstId"`
	WeakAgainstID   int    `json:"weakAgainstId"`
}

type PetTypesData struct {
	PetTypes []PetTypes `json:"petTypes"`
}
