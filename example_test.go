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
		log.Fatalf("could not get teams: %s", err.Error())
	}

	fmt.Printf("Team 1 name: %s, Team 2 name: %s\n", teams[0].Name, teams[1].Name)
	// Output: Team 1 name: New Jersey Devils, Team 2 name: New York Islanders
}

func ExampleGetAllTeams() {
	h := &http.Client{}
	nhlClient := gonhl.NewClient(h)
	nhlClient.Locale = "en_US"

	// team ids 1,2 are NJ Devils and NY Islanders
	teams, err := nhlClient.GetAllTeams()
	if err != nil {
		log.Fatalf("could not get teams: %s", err.Error())
	}

	fmt.Printf("Team 3 name: %s\n", teams[2].Name)
	// Output: Team 3 name: New York Rangers
}
