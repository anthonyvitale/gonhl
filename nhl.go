package gonhl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const baseURL = "https://statsapi.web.nhl.com/api/v1"

type Client struct {
	httpClient *http.Client
	baseURL    string

	// Locale changes the language of certain response values.
	// Supported locales can be found here:
	// https://statsapi.web.nhl.com/api/v1/languages
	Locale string
}

// Client is a client for working with the NHL API.
func NewClient(h *http.Client) *Client {
	return &Client{
		httpClient: h,
		baseURL:    baseURL,
	}
}

func (c *Client) get(uri string, v interface{}) error {
	if c.Locale != "" {
		uri = buildURI(uri, map[string]string{"locale": c.Locale})
	}

	req, err := http.NewRequest("GET", uri, nil)
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

func buildURI(uri string, queryParams map[string]string) string {
	u, _ := url.Parse(uri)
	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
