package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/munsy/gobattlenet/client/sc2"
)

var (
	sc2Command = flag.NewFlagSet("sc2", flag.ExitOnError)

	sc2IDFlag   = sc2Command.Int("id", 0, "ID for various commands (required)")
	sc2NameFlag = sc2Command.String("name", "", "Name for various profile commands.")

	sc2CharacterFlag     = sc2Command.Bool("character", false, "Gets the Sc2 character profile.")
	sc2LadderSeasonsFlag = sc2Command.Bool("ladder-seasons", false, "Gets the Sc2 profile's ladder seasons.")
	sc2MatchHistoryFlag  = sc2Command.Bool("history", false, "Gets the Sc2 profile's match history.")
	sc2LadderFlag        = sc2Command.Bool("ladder", false, "Gets Sc2 ladder data.")
	sc2AchievementsFlag  = sc2Command.Bool("achievements", false, "Gets Sc2 achievement data.")
	sc2RewardsFlag       = sc2Command.Bool("rewards", false, "Gets Sc2 reward data.")
)

func parseSC2Command() {
	if len(os.Args) < 3 {
		printSc2CommandsAndQuit()
	}

	checkAllGameConfigs()
	client, err := sc2.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	if *sc2IDFlag != 0 {
		if *sc2NameFlag != "" {
			if *sc2CharacterFlag {
				result, err := client.Character(*sc2NameFlag, *sc2IDFlag)

				if nil != err {
					panic(err)
				}

				printResult(result)
			} else if *sc2LadderSeasonsFlag {
				result, err := client.LadderSeasons(*sc2NameFlag, *sc2IDFlag)

				if nil != err {
					panic(err)
				}

				printResult(result)
			} else if *sc2MatchHistoryFlag {
				result, err := client.MatchHistory(*sc2NameFlag, *sc2IDFlag)

				if nil != err {
					panic(err)
				}

				printResult(result)
			}
		} else {
			if *sc2LadderFlag {
				result, err := client.Ladder(*sc2IDFlag)

				if nil != err {
					panic(err)
				}

				printResult(result)
			} else if *sc2AchievementsFlag {
				result, err := client.Achievements(*sc2IDFlag)

				if nil != err {
					panic(err)
				}

				printResult(result)
			} else if *sc2RewardsFlag {
				result, err := client.Rewards(*sc2IDFlag)

				if nil != err {
					panic(err)
				}

				printResult(result)
			}
		}
	} else {
		printSc2CommandsAndQuit()
	}
}

func printSc2CommandsAndQuit() {
	if len(os.Args) < 3 {
		fmt.Println("[ERROR] No arguments supplied.")
	} else {
		fmt.Print("[ERROR] Invalid argument. ")
	}
	fmt.Println("Possible arguments are:")
	sc2Command.PrintDefaults()
	os.Exit(1)
}
