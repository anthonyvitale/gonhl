package nhl

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/fortytw2/leaktest"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type NHLSuite struct {
	suite.Suite
	*require.Assertions
	ctx        context.Context
	testServer *httptest.Server
	client     *client
}

func (suite *NHLSuite) SetupTest() {
	suite.Assertions = suite.Suite.Require()
	suite.ctx = context.Background()

	log.SetFormatter(&log.JSONFormatter{})

	suite.testServer = httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do nothing -- should be overwritten in each test:
		// testServer.Config.Handler = <test-specific-handler>
		// and then start the server with `testServer.StartTLS()` or Start()
		// and then set the client.httpClient = testServer.Client()
	}))

	suite.client = &client{}
}

func (suite *NHLSuite) TearDownTest() {
	suite.testServer.Close()
}

func TestNHLSuite(t *testing.T) {
	defer leaktest.Check(t)
	suite.Run(t, new(NHLSuite))
}

func (suite *NHLSuite) Test_joinIntIDs() {
	type args struct {
		ids []int
		sep string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{ids: []int{}, sep: ","}, want: ""},
		{args: args{ids: []int{1}, sep: ","}, want: "1"},
		{args: args{ids: []int{1, 2}, sep: ","}, want: "1,2"},
		{args: args{ids: []int{22, 44, 66}, sep: ","}, want: "22,44,66"},
		{args: args{ids: []int{-1, -2, -3, 0}, sep: "Q"}, want: "-1Q-2Q-3Q0"},
	}
	for i, tt := range tests {
		got := joinIntIDs(tt.args.ids, tt.args.sep)
		suite.Equal(tt.want, got, fmt.Sprintf("test #%d", i))
	}
}

//---------------------------------------------------------------------------------------------------------------------
//
// Helpers
//
//---------------------------------------------------------------------------------------------------------------------

func (suite *NHLSuite) initServer(handler http.Handler) {
	suite.testServer.Config.Handler = handler
	suite.testServer.StartTLS()

	suite.client.httpClient = suite.testServer.Client()

	url, err := url.Parse(suite.testServer.URL)
	suite.NoError(err)

	suite.client.host = url
}
