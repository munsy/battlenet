package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/munsy/gobattlenet/client/d3"
)

var (
	d3Command            = flag.NewFlagSet("d3", flag.ExitOnError)
	d3ActIndexFlag       = d3Command.Bool("actindex", false, "Get a list of all acts (no arguments).")
	d3ActFlag            = d3Command.Int("act", 0, "Get act by ID.")
	d3ArtisanFlag        = d3Command.String("artisan", "", "Get a single artisan by slug.")
	d3RecipeFlag         = d3Command.String("recipe", "", "Get a single recipe by slug for the specified artisan.")
	d3FollowerFlag       = d3Command.String("follower", "", "Get a single follower by slug.")
	d3CharacterClassFlag = d3Command.String("cc", "", "Get a single character class by slug.")
	d3CharacterSkillFlag = d3Command.String("cs", "", "Get a single skill by slug, for a specific character class.")
	d3ItemTypeIndexFlag  = d3Command.Bool("iti", false, "Get an index of item types.")
	d3ItemTypeFlag       = d3Command.String("it", "", "Get a single item type by slug.")
	d3ItemFlag           = d3Command.String("item", "", "Get a single item by item slug and ID.")
	d3AccountFlag        = d3Command.String("account", "", "Get a single profile.")
	d3HeroFlag           = d3Command.Int("hero", -1, "Get a single hero.")
	d3HeroItemsFlag      = d3Command.Bool("hi", false, "Get a list of items for the specified hero.")
	d3FollowerItemsFlag  = d3Command.Bool("fi", false, "Get a list of items for the specified hero's followers.")
)

func parseD3Command() {
	if len(os.Args) < 3 {
		printD3CommandsAndQuit()
	}

	checkAllGameConfigs()
	client, err := d3.New(config.Settings("key"))

	if nil != err {
		panic(err)
	}

	if *d3ActIndexFlag {
		response, err := client.ActIndex()

		if nil != err {
			panic(err)
		}
		if nil == response {
			fmt.Println("Nil response received.")
			os.Exit(1)
		}
		printResult(response)
	} else if *d3ActFlag != 0 {
		response, err := client.Act(*d3ActFlag)

		if nil != err {
			panic(err)
		}
		if nil == response {
			fmt.Println("Nil response received.")
			os.Exit(1)
		}
		printResult(response)
	} else if *d3ArtisanFlag != "" {
		if *d3RecipeFlag != "" {
			response, err := client.Recipe(*d3ArtisanFlag, *d3RecipeFlag)

			if nil != err {
				panic(err)
			}
			if nil == response {
				fmt.Println("Nil response received.")
				os.Exit(1)
			}
			printResult(response)
		} else {
			response, err := client.Artisan(*d3ArtisanFlag)

			if nil != err {
				panic(err)
			}
			if nil == response {
				fmt.Println("Nil response received.")
				os.Exit(1)
			}
			printResult(response)
		}
	} else if *d3FollowerFlag != "" {
		response, err := client.Follower(*d3FollowerFlag)

		if nil != err {
			panic(err)
		}
		if nil == response {
			fmt.Println("Nil response received.")
			os.Exit(1)
		}
		printResult(response)
	} else if *d3CharacterClassFlag != "" {
		if *d3CharacterSkillFlag != "" {
			response, err := client.CharacterSkill(*d3CharacterClassFlag, *d3CharacterSkillFlag)

			if nil != err {
				panic(err)
			}
			if nil == response {
				fmt.Println("Nil response received.")
				os.Exit(1)
			}
			printResult(response)
		} else {
			response, err := client.CharacterClass(*d3CharacterClassFlag)

			if nil != err {
				panic(err)
			}
			if nil == response {
				fmt.Println("Nil response received.")
				os.Exit(1)
			}
			printResult(response)
		}
	} else if *d3ItemTypeIndexFlag == true {
		response, err := client.ItemTypeIndex()

		if nil != err {
			panic(err)
		}
		if nil == response {
			fmt.Println("Nil response received.")
			os.Exit(1)
		}
		printResult(response)
	} else if *d3ItemTypeFlag != "" {
		response, err := client.ItemType(*d3ItemTypeFlag)

		if nil != err {
			panic(err)
		}
		if nil == response {
			fmt.Println("Nil response received.")
			os.Exit(1)
		}
		printResult(response)
	} else if *d3ItemFlag != "" {
		response, err := client.Item(*d3ItemFlag)

		if nil != err {
			panic(err)
		}
		if nil == response {
			fmt.Println("Nil response received.")
			os.Exit(1)
		}
		printResult(response)
	} else if *d3AccountFlag != "" {
		if *d3HeroFlag != -1 {
			if *d3HeroItemsFlag == true {
				response, err := client.HeroItems(*d3AccountFlag, *d3HeroFlag)

				if nil != err {
					panic(err)
				}
				if nil == response {
					fmt.Println("Nil response received.")
					os.Exit(1)
				}
				printResult(response)
			} else if *d3FollowerItemsFlag == true {
				response, err := client.FollowerItems(*d3AccountFlag, *d3HeroFlag)

				if nil != err {
					panic(err)
				}
				if nil == response {
					fmt.Println("Nil response received.")
					os.Exit(1)
				}
				printResult(response)
			} else {
				response, err := client.Hero(*d3AccountFlag, *d3HeroFlag)

				if nil != err {
					panic(err)
				}
				if nil == response {
					fmt.Println("Nil response received.")
					os.Exit(1)
				}
				printResult(response)
			}
		} else {
			response, err := client.Account(*d3AccountFlag)

			if nil != err {
				panic(err)
			}
			if nil == response {
				fmt.Println("Nil response received.")
				os.Exit(1)
			}
			printResult(response)
		}
	} else {
		printD3CommandsAndQuit()
	}
}

func printD3CommandsAndQuit() {
	if len(os.Args) < 3 {
		fmt.Println("[ERROR] No arguments supplied.")
	} else {
		fmt.Print("[ERROR] Invalid argument. ")
	}
	fmt.Println("Possible arguments are:")
	d3Command.PrintDefaults()
	os.Exit(1)
}
