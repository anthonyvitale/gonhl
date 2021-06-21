package gonhl

import (
	"net/http"
	"testing"
)

func TestClient_GetAllTeams(t *testing.T) {
	client, server, err := makeFileTCS(http.StatusOK, "testdata/allteams.json")
	if err != nil {
		t.Fatalf("could not open allteams file: %s", err.Error())
	}

	defer server.Close()
	teams, err := client.GetAllTeams()
	if err != nil {
		t.Fatalf("unexpected http error: %s", err.Error())
	}

	if len(teams) != 32 {
		t.Errorf("getAllTeams() length = %d, want = %d", len(teams), 32)
	}
	if teams[31].Name != "Seattle Kraken" {
		t.Errorf("getAllTeams() pos 31 name = %s, want = %s", teams[31].Name, "Seattle Kraken")
	}
}
