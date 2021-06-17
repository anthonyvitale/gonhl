package gonhl

import (
	"fmt"
)

// TimeZone provides specific details about the team's venue/arena time zone.
type TimeZone struct {
	ID     string `json:"id"`
	Offset int    `json:"offset"`
	Tz     string `json:"tz"`
}

// Venue provides details about the team's current venue/arena.
type Venue struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Link     string   `json:"link"`
	City     string   `json:"city"`
	TimeZone TimeZone `json:"timeZone"`
}

// Division provides details about the team's current division.
type Division struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Conference provides details about the team's current conference.
type Conference struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Franchise provides details about the team's franchise.
type Franchise struct {
	FranchiseID int    `json:"franchiseId"`
	TeamName    string `json:"teamName"`
	Link        string `json:"link"`
}

// Team contains all available information about the team.
type Team struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Link            string     `json:"link"`
	Venue           Venue      `json:"venue"`
	Abbreviation    string     `json:"abbreviation"`
	TeamName        string     `json:"teamName"`
	LocationName    string     `json:"locationName"`
	FirstYearOfPlay string     `json:"firstYearOfPlay"`
	Division        Division   `json:"division"`
	Conference      Conference `json:"conference"`
	Franchise       Franchise  `json:"franchise"`
	ShortName       string     `json:"shortName"`
	OfficialSiteURL string     `json:"officialSiteUrl"`
	FranchiseID     int        `json:"franchiseId"`
	Active          bool       `json:"active"`

	// expand fields
	Roster []Roster `json:"roster,omitempty"`
}

func (c *Client) GetTeam(id int) (*Team, error) {
	t, err := c.GetTeams(id)
	if err != nil {
		return nil, err
	}
	return t[0], nil
}

func (c *Client) GetTeams(ids ...int) ([]*Team, error) {
	teamsURL := fmt.Sprintf("%s/teams?teamId=%s", c.baseURL, joinIntIDs(ids, ","))

	var t struct {
		Teams []*Team `json:"Teams"`
	}

	err := c.get(teamsURL, &t)
	if err != nil {
		return nil, err
	}

	return t.Teams, nil
}

// Person
type Person struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Link     string `json:"link"`
}

type Position struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}

type Roster struct {
	Person       Person   `json:"person"`
	JerseyNumber string   `json:"jerseyNumber"`
	Position     Position `json:"position"`
}

func (c *Client) GetTeamRoster(id int) ([]*Roster, error) {
	teamURL := fmt.Sprintf("%s/teams/%d/roster", c.baseURL, id)

	var t struct {
		Roster []*Roster `json:"Roster"`
	}

	err := c.get(teamURL, &t)
	if err != nil {
		return nil, err
	}

	return t.Roster, nil
}
