package mastodon

import (
	"encoding/json"
	"fmt"
)

const (
	TrendsLinksURI    string = "/api/v1/trends/links"
	TrendsStatusesURI string = "/api/v1/trends/statuses"
	TrendsTagsURI     string = "/api/v1/trends/tags"
)

// TrendsLinks hold information on links
type TrendLinks []struct {
	URL          string `json:"url"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	AuthorName   string `json:"author_name"`
	AuthorURL    string `json:"author_url"`
	ProviderName string `json:"provider_name"`
	ProviderURL  string `json:"provider_url"`
	HTML         string `json:"html"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Image        string `json:"image"`
	EmbedURL     string `json:"embed_url"`
	Blurhash     string `json:"blurhash"`
	History      []struct {
		Day      string `json:"day"`
		Accounts string `json:"accounts"`
		Uses     string `json:"uses"`
	} `json:"history"`
}

// TrendsTags hold information for tags
type TrendTags []struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	History []struct {
		Day      string `json:"day"`
		Accounts string `json:"accounts"`
		Uses     string `json:"uses"`
	} `json:"history"`
	Following bool `json:"following"`
}

// Get links that have been shared more than others
func (c *Client) GetTrendsLinks() (TrendLinks, error) {
	links := TrendLinks{}

	url := fmt.Sprintf("https://%s%s", c.Server, TrendsLinksURI)

	body, err := c.SendRequest(url)
	if err != nil {
		return links, err
	}

	err = json.Unmarshal(body, &links)

	return links, err
}

// Get statuses that have been interacted with more than others
// func (c *Client) GetTrendsStatuses() (TrendStatuses, error) {
// 	statuses := TrendStatuses{}

// 	url := fmt.Sprintf("https://%s%s", c.Server, TrendsStatusesURI)

// 	body, err := c.SendRequest(url)
// 	if err != nil {
// 		return statuses, err
// 	}

// 	err = json.Unmarshal(body, &statuses)

// 	return statuses, err
// }

// Get tags that are being used more frequently within the past week
func (c *Client) GetTrendsTags() (TrendTags, error) {
	tags := TrendTags{}

	url := fmt.Sprintf("https://%s%s", c.Server, TrendsTagsURI)

	body, err := c.SendRequest(url)
	if err != nil {
		return tags, err
	}

	err = json.Unmarshal(body, &tags)

	return tags, err
}
