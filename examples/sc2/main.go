package main

import (
	"flag"
	"fmt"

	"github.com/munsy/gobattlenet/sc2"
)

var keyFlag = flag.String("k", "", "Battle.net API key (required).")
var achievementFlag = flag.Int("id", 0, "Starcraft II achievement ID number (required).")

func main() {
	flag.Parse()

	if *achievementFlag == 0 {
		fmt.Println("Invalid achivement id.")
		return
	}

	client, err := sc2.New(*keyFlag)

	if nil != err {
		panic(err)
	}

	achievements, err := client.Achievements(*achievementFlag)

	if nil != err {
		panic(err)
	}

	fmt.Printf("%s\t\t%s\n", "AchievementID", "Title")
	fmt.Println("----------------------------------------")
	for i := 0; i < len(achievements.Achievements); i++ {
		fmt.Printf("%d\t\t%s\n", achievements.Achievements[i].AchievementID, achievements.Achievements[i].Title)
	}
}
