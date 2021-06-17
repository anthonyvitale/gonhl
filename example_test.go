package gonhl_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anthonyvitale/gonhl"
)

func ExampleGetTeams() {
	h := &http.Client{}
	nhlClient := gonhl.NewClient(h)
	nhlClient.Locale = "en_US"

	// team ids 1,2 are NJ Devils and NY Islanders
	teams, err := nhlClient.GetTeams(1, 2)
	if err != nil {
		log.Fatalf("couldn't search for teams: %s", err.Error())
	}

	fmt.Printf("Team 1 name: %s, Team 2 name: %s\n", teams[0].Name, teams[1].Name)
	// Output: Team 1 name: New Jersey Devils, Team 2 name: New York Islanders
}
