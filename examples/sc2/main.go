package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet"
	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/regions"
)

var keyFlag = flag.String("k", "", "Battle.net API key (required).")
var achievementFlag = flag.Int("id", 0, "Starcraft II achievement ID number (required).")

func main() {
	flag.Parse()

	if *keyFlag == "" {
		fmt.Println("No key provided.")
		return
	}

	if *achievementFlag == 0 {
		fmt.Println("Invalid achievement id.")
		return
	}

	settings := &battlenet.Settings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
	}

	client, err := battlenet.SC2Client(settings, *keyFlag)

	if nil != err {
		panic(err)
	}

	response, err := client.Achievements(*achievementFlag)

	if nil != err {
		panic(err)
	}

	achievements := response.Data

	fmt.Printf("%s\t\t%s\n", "AchievementID", "Title")
	fmt.Println("----------------------------------------")
	for i := 0; i < len(achievements.Achievements); i++ {
		fmt.Printf("%d\t\t%s\n", achievements.Achievements[i].AchievementID, achievements.Achievements[i].Title)
	}
}
