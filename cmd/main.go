package main

import (
	"fmt"
	"os"
)

var config Config

func init() {
	readTOML(*configFlag)
	if *keyFlag == "" {
		*keyFlag = config.Key
	}
	if *tokenFlag == "" {
		*tokenFlag = config.Token
	}
	if *regionFlag == "" {
		*regionFlag = config.Region
	}
	if *localeFlag == "" {
		*localeFlag = config.Locale
	}
}

func main() {
	if len(os.Args) < 2 {
		printDefaultsAndQuit()
	}

	switch os.Args[1] {
	case "config":
		configCommand.Parse(os.Args[2:])
	case "account":
		accountCommand.Parse(os.Args[2:])
	case "d3":
		d3Command.Parse(os.Args[2:])
	case "sc2":
		sc2Command.Parse(os.Args[2:])
	case "wow":
		wowCommand.Parse(os.Args[2:])
	default:
		printDefaultsAndQuit()
	}

	if configCommand.Parsed() {
		parseConfigCommand()
	}
	if accountCommand.Parsed() {
		parseAccountCommand()
	}
	if d3Command.Parsed() {
		parseD3Command()
	}
	if sc2Command.Parsed() {
		parseSC2Command()
	}
	if wowCommand.Parsed() {
		parseWoWCommand()
	}
}

func printDefaultsAndQuit() {
	if len(os.Args) < 2 {
		fmt.Print("[ERROR] No arguments supplied. ")
	} else {
		fmt.Print("[ERROR] Invalid argument. ")
	}
	fmt.Println("Possible arguments are:")
	fmt.Println("\t- config (Configuration settings)")
	fmt.Println("\t- account (Account API access)")
	fmt.Println("\t- d3 (Diablo III API access)")
	fmt.Println("\t- sc2 (Starcraft II API access)")
	fmt.Println("\t- wow (World of Warcraft API access)")
	os.Exit(1)
}
