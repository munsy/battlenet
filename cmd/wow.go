package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/munsy/gobattlenet/client/wow"
)

var (
	wowCommand = flag.NewFlagSet("wow", flag.ExitOnError)

	wowIDFlag            = wowCommand.Int("id", 0, "ID for various commands (required)")
	wowCharacterNameFlag = wowCommand.String("cn", "", "Character name for various commands.")
	wowRealmNameFlag     = wowCommand.String("rn", "", "Realm name for various commands.")
	wowGuildNameFlag     = wowCommand.String("gn", "", "Guild name for various commands.")

	wowAchievementFlag           = wowCommand.Bool("achievement", false, "Gets data about an individual achievement.")
	wowAuctionFlag               = wowCommand.Bool("auction", false, "Gets rolling batches of data about current auctions.")
	wowBossMasterListFlag        = wowCommand.Bool("boss-master-list", false, "Gets a list of all supported bosses.")
	wowBossFlag                  = wowCommand.Bool("boss", false, "Gets information about a boss.")
	wowCharacterProfileFlag      = wowCommand.Bool("character-profile", false, "Gets character information.")
	wowGuildProfileFlag          = wowCommand.Bool("guild-profile", false, "Gets guild information.")
	wowItemFlag                  = wowCommand.Bool("item", false, "Gets detailed item information.")
	wowItemSetFlag               = wowCommand.Bool("item-set", false, "Gets detailed item set information.")
	wowMountMasterListFlag       = wowCommand.Bool("mount-master-list", false, "Gets a list of all supported mounts.")
	wowPetMasterListFlag         = wowCommand.Bool("petMaster-list", false, "Gets a list of all supported battle and vanity pets.")
	wowPetAbilitiesFlag          = wowCommand.Bool("pet-abilities", false, "Gets data about an individual battle pet ability ID.")
	wowPetSpeciesFlag            = wowCommand.Bool("pet-species", false, "Gets data about an individual pet species.")
	wowPetStatsFlag              = wowCommand.Bool("pet-stats", false, "Gets detailed information about the given species of pet.")
	wowQuestFlag                 = wowCommand.Bool("quest", false, "Gets metadata for a given quest.")
	wowRealmStatusFlag           = wowCommand.Bool("realm-status", false, "Gets realm status information.")
	wowRecipeFlag                = wowCommand.Bool("recipe", false, "Gets basic recipe information.")
	wowSpellFlag                 = wowCommand.Bool("spell", false, "Gets information about spells.")
	wowZoneMasterListFlag        = wowCommand.Bool("zone-master-list", false, "Gets a list of all supported zones and their bosses.")
	wowZoneFlag                  = wowCommand.Bool("zone", false, "Gets information about zones.")
	wowBattlegroupsFlag          = wowCommand.Bool("battlegroups", false, "Gets a list of battlegroups for this region.")
	wowCharacterRacesFlag        = wowCommand.Bool("character-races", false, "Gets a list of each race and their associated faction, name, unique ID, and skin.")
	wowCharacterClassesFlag      = wowCommand.Bool("character-classes", false, "Gets a list of character classes.")
	wowCharacterAchievementsFlag = wowCommand.Bool("character-achievements", false, "Gets a list of all of the achievements that characters can earn as well as the category structure and hierarchy.")
	wowGuildRewardsFlag          = wowCommand.Bool("guild-rewards", false, "Gets a list of all guild rewards.")
	wowGuildPerksFlag            = wowCommand.Bool("guild-perks", false, "Gets a list of all guild perks.")
	wowGuildAchievementsFlag     = wowCommand.Bool("guild-achievements", false, "Gets a list of all of the achievements that guilds can earn as well as the category structure and hierarchy.")
	wowItemClassesFlag           = wowCommand.Bool("item-classes", false, "Gets a list of item classes.")
	wowTalentsFlag               = wowCommand.Bool("talents", false, "Gets a list of talents, specs and glyphs for each class.")
	wowPetTypesFlag              = wowCommand.Bool("pet-types", false, "Gets a list of different pet types (including what they are strong and weak against).")
)

func parseWoWCommand() {
	checkAllGameConfigs()
	client, err := wow.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	if *wowIDFlag != 0 {
		if *wowAchievementFlag == true {
			result, err := client.Achievement(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowItemFlag == true {
			result, err := client.Item(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowItemSetFlag == true {
			result, err := client.ItemSet(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowBossFlag == true {
			result, err := client.Boss(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowPetAbilitiesFlag == true {
			result, err := client.PetAbilities(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowPetSpeciesFlag == true {
			result, err := client.PetSpecies(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowPetStatsFlag == true {
			result, err := client.PetStats(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowQuestFlag == true {
			result, err := client.Quest(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowRecipeFlag == true {
			result, err := client.Recipe(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowSpellFlag == true {
			result, err := client.Spell(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowZoneFlag == true {
			result, err := client.Zone(*wowIDFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		}
	} else if *wowRealmNameFlag != "" {
		if *wowCharacterNameFlag != "" {
			result, err := client.CharacterProfile(*wowRealmNameFlag, *wowCharacterNameFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowGuildNameFlag != "" {
			result, err := client.GuildProfile(*wowRealmNameFlag, *wowGuildNameFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		} else if *wowAuctionFlag == true {
			result, err := client.AuctionDataStatus(*wowRealmNameFlag)

			if nil != err {
				panic(err)
			}

			printResult(result)
		}
	} else if *wowBossMasterListFlag == true {
		result, err := client.BossMasterList()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowMountMasterListFlag == true {
		result, err := client.MountMasterList()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowPetMasterListFlag == true {
		result, err := client.PetMasterList()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowRealmStatusFlag == true {
		result, err := client.RealmStatus()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowZoneMasterListFlag == true {
		result, err := client.ZoneMasterList()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowBattlegroupsFlag == true {
		result, err := client.Battlegroups()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowCharacterRacesFlag == true {
		result, err := client.CharacterRaces()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowCharacterClassesFlag == true {
		result, err := client.CharacterClasses()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowCharacterAchievementsFlag == true {
		result, err := client.CharacterAchievements()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowGuildRewardsFlag == true {
		result, err := client.GuildRewards()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowGuildPerksFlag == true {
		result, err := client.GuildPerks()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowGuildAchievementsFlag == true {
		result, err := client.GuildAchievements()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowItemClassesFlag == true {
		result, err := client.ItemClasses()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowTalentsFlag == true {
		result, err := client.Talents()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else if *wowPetTypesFlag == true {
		result, err := client.PetTypes()

		if nil != err {
			panic(err)
		}

		printResult(result)
	} else {
		printWoWCommandsAndQuit()
	}

}

func printWoWCommandsAndQuit() {
	if len(os.Args) < 3 {
		fmt.Println("[ERROR] No arguments supplied.")
	} else {
		fmt.Print("[ERROR] Invalid argument. ")
	}
	fmt.Println("Possible arguments are:")
	wowCommand.PrintDefaults()
	os.Exit(1)
}
