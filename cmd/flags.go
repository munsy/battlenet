package main

import (
	"flag"
)

var (
	// Subcommands
	configCommand  = flag.NewFlagSet("config", flag.ExitOnError)
	accountCommand = flag.NewFlagSet("account", flag.ExitOnError)
	d3Command      = flag.NewFlagSet("d3", flag.ExitOnError)
	sc2Command     = flag.NewFlagSet("sc2", flag.ExitOnError)
	wowCommand     = flag.NewFlagSet("wow", flag.ExitOnError)

	// Config flags
	configFlag = flag.String("f", "config.toml", "Configuration file to load key/token values.")
	keyFlag    = flag.String("k", "", "Set the Battle.net client key for API access.")
	tokenFlag  = flag.String("t", "", "Set the OAuth2 token for Battle.net profile API access.")
	writeFlag  = flag.Bool("o", false, "Write current key/token settings in memory to config file.")
	regionFlag = flag.String("r", "", "Set the Battle.net region.") // ADD
	localeFlag = flag.String("l", "", "Set the Battle.net locale.") // ADD

	// Account flags
	accountBattleIDFlag   = flag.Bool("bid", false, "Display Battle.net ID and BattleTag.")
	accountSc2ProfileFlag = flag.Bool("sc2", false, "Display details about a Starcraft II character.")
	accountWoWProfileFlag = flag.Bool("wow", false, "Display details about a World of Warcraft profile.")

/*
	// D3 flags
	d3Flag = flag.Bool("", , "")
	d3Flag = flag.Bool("", , "")
	d3Flag = flag.Bool("", , "")
	d3Flag = flag.Bool("", , "")
	d3Flag = flag.Bool("", , "")
	d3Flag = flag.Bool("", , "")
	d3Flag = flag.Bool("", , "")

	// SC2 flags
	sc2Flag = flag.Bool("", , "")
	sc2Flag = flag.Bool("", , "")
	sc2Flag = flag.Bool("", , "")
	sc2Flag = flag.Bool("", , "")
	sc2Flag = flag.Bool("", , "")
	sc2Flag = flag.Bool("", , "")
	sc2Flag = flag.Bool("", , "")

	// WoW flags
	wowFlag = flag.Bool("", , "")
	wowFlag = flag.Bool("", , "")
	wowFlag = flag.Bool("", , "")
	wowFlag = flag.Bool("", , "")
	wowFlag = flag.Bool("", , "")
	wowFlag = flag.Bool("", , "")
	wowFlag = flag.Bool("", , "")
*/
)
