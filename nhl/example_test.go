package nhl_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anthonyvitale/gonhl/nhl"
)

const baseURL = "https://statsapi.web.nhl.com/api/v1"

func ExampleGetTeams() {
	client, err := nhl.NewClient(http.DefaultClient, baseURL)
	if err != nil {
		log.Fatalf("could not create client: %s", err.Error())
	}

	// team ids 1,2 are NJ Devils and NY Islanders
	teams, err := client.GetTeams(1, 2)
	if err != nil {
		log.Fatalf("could not get teams: %s", err.Error())
	}

	fmt.Printf("Team 1 name: %s, Team 2 name: %s\n", teams[0].Name, teams[1].Name)
	// Output: Team 1 name: New Jersey Devils, Team 2 name: New York Islanders
}

func ExampleGetAllTeams() {
	client, err := nhl.NewClient(http.DefaultClient, baseURL)
	if err != nil {
		log.Fatalf("could not create client: %s", err.Error())
	}

	// team ids 1,2 are NJ Devils and NY Islanders
	teams, err := client.GetAllTeams()
	if err != nil {
		log.Fatalf("could not get teams: %s", err.Error())
	}

	fmt.Printf("Team 3 name: %s\n", teams[2].Name)
	// Output: Team 3 name: New York Rangers
}
