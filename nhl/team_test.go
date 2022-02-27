package nhl

import (
	"fmt"
	"net/http"
	"os"
)

func (suite *NHLSuite) TestGetAllTeams() {

	suite.initServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("testdata/allteams.json")
		suite.NoError(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(data))
	}))

	teams, err := suite.client.GetAllTeams()
	suite.NoError(err)
	suite.Len(teams, 32)
	// Spot check
	suite.Equal("Seattle Kraken", teams[31].Name)
}
