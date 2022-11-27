package mastodon

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Convenience constants Mastodon
const (
	UserAgent = "mastodon-public-api"
	Timeout   = 5
)

// Client is a API client for mastodon.
type Client struct {
	http.Client
	Server    string
	UserAgent string
}

// NewClient returns a new mastodon API client.
func NewClient(server string) (*Client, error) {
	// Check that the user provided a valid schema
	serverlower := strings.ToLower(server)
	if !strings.HasPrefix(serverlower, "https://") && !strings.HasPrefix(serverlower, "http://") {
		e := fmt.Sprintf("invalid server provided: %s", server)
		return &Client{}, errors.New(e)
	}

	c := &Client{
		Client:    *http.DefaultClient,
		Server:    server,
		UserAgent: UserAgent,
	}

	// Set default timeout
	c.Client.Timeout = time.Second * Timeout

	return c, nil
}

// Send request and obtain body
func (c *Client) SendRequest(url string) ([]byte, error) {
	var data []byte

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}

	req.Header.Set("User-Agent", c.UserAgent)

	// Send request
	resp, err := c.Client.Do(req)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	// Verify response was 200
	if resp.StatusCode != 200 {
		err = errors.New(
			"resp.StatusCode: " +
				strconv.Itoa(resp.StatusCode))
		return data, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	data = body

	return data, nil
}
