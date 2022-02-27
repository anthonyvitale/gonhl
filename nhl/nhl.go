package nhl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

const baseURL = "https://statsapi.web.nhl.com/api/v1"

// Client is a client for working with the NHL API.
type Client interface {
	GetTeam(int) (*Team, error)
	GetTeams(...int) ([]*Team, error)
	GetAllTeams() ([]*Team, error)
	GetTeamRoster(int) ([]*Roster, error)
}

type client struct {
	httpClient *http.Client
	host       *url.URL

	// Locale changes the language of certain response values.
	// Supported locales can be found here:
	// https://statsapi.web.nhl.com/api/v1/languages
	locale string
}

// NewClient is a client for working with the NHL API.
func NewClient(httpClient *http.Client, host string) (Client, error) {
	if host == "" {
		return nil, fmt.Errorf("nhl: host cannot be empty")
	}

	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	if httpClient == nil {
		return nil, fmt.Errorf("nhl: http client cannot be nil")
	}

	return &client{
		httpClient: httpClient,
		host:       u,
	}, nil
}

func (c *client) get(url string, v interface{}) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("gonhl: error HTTP %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}

func joinIntIDs(ids []int, sep string) string {
	var sb strings.Builder
	for i, v := range ids {
		sb.WriteString(strconv.Itoa(v))
		if i < len(ids)-1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}

func buildURL(u *url.URL, endpoint string, queryParams map[string]string) string {
	u.Path = path.Join(u.Path, endpoint)
	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
