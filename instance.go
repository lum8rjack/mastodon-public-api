package mastodon

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	InstanceURI                string = "/api/v2/instance"
	InstanceActivityURI        string = "/api/v1/instance/activity"
	InstanceDomainsBlockedyURI string = "/api/v1/instance/domain_block"
	InstancePeersURI           string = "/api/v1/instance/peers"
	InstanceRulesURI           string = "/api/v1/instance/rules"
)

// Instance hold information for instance
type Instance struct {
	Domain      string `json:"domain"`
	Title       string `json:"title"`
	Version     string `json:"version"`
	SourceURL   string `json:"source_url"`
	Description string `json:"description"`
	Usage       struct {
		Users struct {
			ActiveMonth int `json:"active_month"`
		} `json:"users"`
	} `json:"usage"`
	Thumbnail struct {
		URL      string      `json:"url"`
		Blurhash interface{} `json:"blurhash"`
		Versions struct {
			One_X string `json:"@1x"`
			Two_X string `json:"@2x"`
		} `json:"versions"`
	} `json:"thumbnail"`
	Languages     []string `json:"languages"`
	Configuration struct {
		Urls struct {
			Streaming string `json:"streaming"`
		} `json:"urls"`
		Accounts struct {
			MaxFeaturedTags int `json:"max_featured_tags"`
		} `json:"accounts"`
		Statuses struct {
			MaxCharacters            int `json:"max_characters"`
			MaxMediaAttachments      int `json:"max_media_attachments"`
			CharactersReservedPerURL int `json:"characters_reserved_per_url"`
		} `json:"statuses"`
		MediaAttachments struct {
			SupportedMimeTypes  []string `json:"supported_mime_types"`
			ImageSizeLimit      int      `json:"image_size_limit"`
			ImageMatrixLimit    int      `json:"image_matrix_limit"`
			VideoSizeLimit      int      `json:"video_size_limit"`
			VideoFrameRateLimit int      `json:"video_frame_rate_limit"`
			VideoMatrixLimit    int      `json:"video_matrix_limit"`
		} `json:"media_attachments"`
		Polls struct {
			MaxOptions             int `json:"max_options"`
			MaxCharactersPerOption int `json:"max_characters_per_option"`
			MinExpiration          int `json:"min_expiration"`
			MaxExpiration          int `json:"max_expiration"`
		} `json:"polls"`
		Translation struct {
			Enabled bool `json:"enabled"`
		} `json:"translation"`
	} `json:"configuration"`
	Registrations struct {
		Enabled          bool        `json:"enabled"`
		ApprovalRequired bool        `json:"approval_required"`
		Message          interface{} `json:"message"`
	} `json:"registrations"`
	Contact struct {
		Email   string `json:"email"`
		Account struct {
			ID             string    `json:"id"`
			Username       string    `json:"username"`
			Acct           string    `json:"acct"`
			DisplayName    string    `json:"display_name"`
			Locked         bool      `json:"locked"`
			Bot            bool      `json:"bot"`
			Discoverable   bool      `json:"discoverable"`
			Group          bool      `json:"group"`
			CreatedAt      time.Time `json:"created_at"`
			Note           string    `json:"note"`
			URL            string    `json:"url"`
			Avatar         string    `json:"avatar"`
			AvatarStatic   string    `json:"avatar_static"`
			Header         string    `json:"header"`
			HeaderStatic   string    `json:"header_static"`
			FollowersCount int       `json:"followers_count"`
			FollowingCount int       `json:"following_count"`
			StatusesCount  int       `json:"statuses_count"`
			LastStatusAt   string    `json:"last_status_at"`
			Noindex        bool      `json:"noindex"`
			Emojis         []struct {
				Shortcode       string `json:"shortcode"`
				URL             string `json:"url"`
				StaticURL       string `json:"static_url"`
				VisibleInPicker bool   `json:"visible_in_picker"`
			} `json:"emojis"`
			Fields []struct {
				Name       string      `json:"name"`
				Value      string      `json:"value"`
				VerifiedAt interface{} `json:"verified_at"`
			} `json:"fields"`
		} `json:"account"`
	} `json:"contact"`
	Rules InstanceRules `json:"rules"`
}

// Rules hold rules for the instance
type InstanceRules []struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// InstancePeers hold information for instance peers
type InstancePeers []string

// InstanceActivity hold information for instance activity
type InstanceActivity []struct {
	Week          string `json:"week"`
	Statuses      string `json:"statuses"`
	Logins        string `json:"logins"`
	Registrations string `json:"registrations"`
}

// DomainsBlocked hold information on domains blocked
type DomainsBlocked []struct {
	Domain   string `json:"domain"`
	Digest   string `json:"digest"`
	Severity string `json:"severity"`
	Comment  string `json:"comment"`
}

// Error for unauthorized requests
type Unauthorized struct {
	Error string `json:"error"`
}

// Get general information about the server
func (c *Client) GetInstanceData() (Instance, error) {
	instance := Instance{}

	url := fmt.Sprintf("%s%s", c.Server, InstanceURI)

	body, err := c.SendRequest(url)
	if err != nil {
		return instance, err
	}

	err = json.Unmarshal(body, &instance)

	return instance, err
}

// Get domains that this instance is aware of
// FIXME: Need to add additional check for if the server is in
// whitelist mode and the Authorization header is missing or invalid
// https://docs.joinmastodon.org/methods/instance/#peers
func (c *Client) GetInstancePeers() (InstancePeers, error) {
	instancepeers := InstancePeers{}

	url := fmt.Sprintf("%s%s", c.Server, InstancePeersURI)

	body, err := c.SendRequest(url)
	if err != nil {
		return instancepeers, err
	}

	err = json.Unmarshal(body, &instancepeers)

	return instancepeers, err
}

// Get instance activity over the last 3 months, binned weekly
// FIXME: Need to add additional check for if the server is in
// whitelist mode and the Authorization header is missing or invalid
// https://docs.joinmastodon.org/methods/instance/#activity
func (c *Client) GetInstanceActivity() (InstanceActivity, error) {
	instanceactivity := InstanceActivity{}

	url := fmt.Sprintf("%s%s", c.Server, InstanceActivityURI)

	body, err := c.SendRequest(url)
	if err != nil {
		return instanceactivity, err
	}

	err = json.Unmarshal(body, &instanceactivity)

	return instanceactivity, err
}

// Get instance rules that the users of this service should follow
func (c *Client) GetInstanceRules() (InstanceRules, error) {
	instancerules := InstanceRules{}

	url := fmt.Sprintf("%s%s", c.Server, InstanceRulesURI)

	body, err := c.SendRequest(url)
	if err != nil {
		return instancerules, err
	}

	err = json.Unmarshal(body, &instancerules)

	return instancerules, err
}

// Get a list of domains that have been blocked
// FIXME: Need to add additional check for if the server is in
// whitelist mode and the Authorization header is missing or invalid
// https://docs.joinmastodon.org/methods/instance/#domain_blocks
func (c *Client) GetInstanceDomainsBlocked() (DomainsBlocked, error) {
	domainsblocked := DomainsBlocked{}

	url := fmt.Sprintf("%s%s", c.Server, InstanceDomainsBlockedyURI)

	body, err := c.SendRequest(url)
	if err != nil {
		return domainsblocked, err
	}

	err = json.Unmarshal(body, &domainsblocked)

	return domainsblocked, err
}
