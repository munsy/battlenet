package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/client/sc2"
	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

var keyFlag = flag.String("k", "", "Battle.net API key (required).")
var achievementFlag = flag.Int("id", 0, "Starcraft II achievement ID number (required).")

func main() {
	flag.Parse()

	if *achievementFlag == 0 {
		fmt.Println("Invalid achievement id.")
		return
	}

	settings := &settings.BNetSettings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
	}

	client, err := http.New(settings)

	if nil != err {
		panic(err)
	}

	achievements, err := client.Sc2(*keyFlag).Achievements(*achievementFlag)

	if nil != err {
		panic(err)
	}

	fmt.Printf("%s\t\t%s\n", "AchievementID", "Title")
	fmt.Println("----------------------------------------")
	for i := 0; i < len(achievements.Achievements); i++ {
		fmt.Printf("%d\t\t%s\n", achievements.Achievements[i].AchievementID, achievements.Achievements[i].Title)
	}
}
