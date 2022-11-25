package mastodon

import (
	"encoding/json"
	"fmt"
)

const (
	CustomEmojisURI string = "/api/v1/custom_emojis"
)

// Emojis hold information for custom emojis
type Emojis []struct {
	Shortcode       string `json:"shortcode"`
	URL             string `json:"url"`
	StaticURL       string `json:"static_url"`
	VisibleInPicker bool   `json:"visible_in_picker"`
}

// Get custom emojis that are available on the server
func (c *Client) GetCustomEmojis() (Emojis, error) {
	customemojis := Emojis{}

	url := fmt.Sprintf("https://%s%s", c.Server, CustomEmojisURI)

	body, err := c.SendRequest(url)
	if err != nil {
		return customemojis, err
	}

	err = json.Unmarshal(body, &customemojis)

	return customemojis, err
}
