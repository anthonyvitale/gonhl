package gonhl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const baseURL = "https://statsapi.web.nhl.com/api/v1"

type Client struct {
	httpClient *http.Client
	baseURL    string

	// Locale changes the language of certain response values.
	// Currently supported locales are as follows:
	// en_US // fr_CA // es_ES // cs_CS // sv_SV // sk_SK // de_DE // ru_RU // fi_FI
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
		u, _ := url.Parse(uri)
		q := u.Query()
		q.Set("locale", c.Locale)
		u.RawQuery = q.Encode()
		uri = u.String()
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
