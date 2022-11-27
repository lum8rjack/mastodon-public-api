package mastodon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testunauthorized string = `{"error": "This API requires an authenticated user"}`
	testinstance     string = `{
		"domain": "mastodon.social",
		"title": "Mastodon",
		"version": "4.0.0rc1",
		"source_url": "https://github.com/mastodon/mastodon",
		"description": "The original server operated by the Mastodon gGmbH non-profit",
		"usage": {
		  "users": {
			"active_month": 123122
		  }
		},
		"thumbnail": {
		  "url": "https://files.mastodon.social/site_uploads/files/000/000/001/@1x/57c12f441d083cde.png",
		  "blurhash": "UeKUpFxuo~R%0nW;WCnhF6RjaJt757oJodS$",
		  "versions": {
			"@1x": "https://files.mastodon.social/site_uploads/files/000/000/001/@1x/57c12f441d083cde.png",
			"@2x": "https://files.mastodon.social/site_uploads/files/000/000/001/@2x/57c12f441d083cde.png"
		  }
		},
		"languages": [
		  "en"
		],
		"configuration": {
		  "urls": {
			"streaming": "wss://mastodon.social"
		  },
		  "accounts": {
			"max_featured_tags": 10
		  },
		  "statuses": {
			"max_characters": 500,
			"max_media_attachments": 4,
			"characters_reserved_per_url": 23
		  },
		  "media_attachments": {
			"supported_mime_types": [
			  "image/jpeg",
			  "image/png",
			  "image/gif",
			  "image/heic",
			  "image/heif",
			  "image/webp",
			  "video/webm",
			  "video/mp4",
			  "video/quicktime",
			  "video/ogg",
			  "audio/wave",
			  "audio/wav",
			  "audio/x-wav",
			  "audio/x-pn-wave",
			  "audio/vnd.wave",
			  "audio/ogg",
			  "audio/vorbis",
			  "audio/mpeg",
			  "audio/mp3",
			  "audio/webm",
			  "audio/flac",
			  "audio/aac",
			  "audio/m4a",
			  "audio/x-m4a",
			  "audio/mp4",
			  "audio/3gpp",
			  "video/x-ms-asf"
			],
			"image_size_limit": 10485760,
			"image_matrix_limit": 16777216,
			"video_size_limit": 41943040,
			"video_frame_rate_limit": 60,
			"video_matrix_limit": 2304000
		  },
		  "polls": {
			"max_options": 4,
			"max_characters_per_option": 50,
			"min_expiration": 300,
			"max_expiration": 2629746
		  },
		  "translation": {
			"enabled": true
		  }
		},
		"registrations": {
		  "enabled": false,
		  "approval_required": false,
		  "message": null
		},
		"contact": {
		  "email": "staff@mastodon.social",
		  "account": {
			"id": "1",
			"username": "Gargron",
			"acct": "Gargron",
			"display_name": "Eugen 💀",
			"locked": false,
			"bot": false,
			"discoverable": true,
			"group": false,
			"created_at": "2016-03-16T00:00:00.000Z",
			"note": "<p>Founder, CEO and lead developer <span class=\"h-card\"><a href=\"https://mastodon.social/@Mastodon\" class=\"u-url mention\">@<span>Mastodon</span></a></span>, Germany.</p>",
			"url": "https://mastodon.social/@Gargron",
			"avatar": "https://files.mastodon.social/accounts/avatars/000/000/001/original/dc4286ceb8fab734.jpg",
			"avatar_static": "https://files.mastodon.social/accounts/avatars/000/000/001/original/dc4286ceb8fab734.jpg",
			"header": "https://files.mastodon.social/accounts/headers/000/000/001/original/3b91c9965d00888b.jpeg",
			"header_static": "https://files.mastodon.social/accounts/headers/000/000/001/original/3b91c9965d00888b.jpeg",
			"followers_count": 133026,
			"following_count": 311,
			"statuses_count": 72605,
			"last_status_at": "2022-10-31",
			"noindex": false,
			"emojis": [],
			"fields": [
			  {
				"name": "Patreon",
				"value": "<a href=\"https://www.patreon.com/mastodon\" target=\"_blank\" rel=\"nofollow noopener noreferrer me\"><span class=\"invisible\">https://www.</span><span class=\"\">patreon.com/mastodon</span><span class=\"invisible\"></span></a>",
				"verified_at": null
			  }
			]
		  }
		},
		"rules": [
		  {
			"id": "1",
			"text": "Sexually explicit or violent media must be marked as sensitive when posting"
		  },
		  {
			"id": "2",
			"text": "No racism, sexism, homophobia, transphobia, xenophobia, or casteism"
		  },
		  {
			"id": "3",
			"text": "No incitement of violence or promotion of violent ideologies"
		  },
		  {
			"id": "4",
			"text": "No harassment, dogpiling or doxxing of other users"
		  },
		  {
			"id": "5",
			"text": "No content illegal in Germany"
		  },
		  {
			"id": "7",
			"text": "Do not share intentionally false or misleading information"
		  }
		]
	  }`

	testinstanceactivity = `[
		{
		  "week": "1574640000",
		  "statuses": "37125",
		  "logins": "14239",
		  "registrations": "542"
		},
		{
		  "week": "1574035200",
		  "statuses": "244447",
		  "logins": "28820",
		  "registrations": "4425"
		},
		{
		  "week": "1573430400",
		  "statuses": "270615",
		  "logins": "35388",
		  "registrations": "8781"
		},
		{
		  "week": "1572825600",
		  "statuses": "309722",
		  "logins": "44433",
		  "registrations": "26165"
		},
		{
		  "week": "1572220800",
		  "statuses": "116227",
		  "logins": "19739",
		  "registrations": "2926"
		},
		{
		  "week": "1571616000",
		  "statuses": "119932",
		  "logins": "19247",
		  "registrations": "3188"
		},
		{
		  "week": "1571011200",
		  "statuses": "117892",
		  "logins": "19164",
		  "registrations": "3107"
		},
		{
		  "week": "1570406400",
		  "statuses": "109092",
		  "logins": "18763",
		  "registrations": "2986"
		},
		{
		  "week": "1569801600",
		  "statuses": "107554",
		  "logins": "19614",
		  "registrations": "2904"
		},
		{
		  "week": "1569196800",
		  "statuses": "118067",
		  "logins": "19703",
		  "registrations": "3295"
		},
		{
		  "week": "1568592000",
		  "statuses": "110199",
		  "logins": "19791",
		  "registrations": "3026"
		},
		{
		  "week": "1567987200",
		  "statuses": "106029",
		  "logins": "19089",
		  "registrations": "2769"
		}
	  ]`

	testinstancerules = `[
		{
		  "id": "1",
		  "text": "Sexually explicit or violent media must be marked as sensitive when posting"
		},
		{
		  "id": "2",
		  "text": "No racism, sexism, homophobia, transphobia, xenophobia, or casteism"
		},
		{
		  "id": "3",
		  "text": "No incitement of violence or promotion of violent ideologies"
		},
		{
		  "id": "4",
		  "text": "No harassment, dogpiling or doxxing of other users"
		},
		{
		  "id": "5",
		  "text": "No content illegal in Germany"
		},
		{
		  "id": "7",
		  "text": "Do not share intentionally false or misleading information"
		}
	  ]`

	testinstancedomainsblocked = `[
		{
		  "domain":"birb.elfenban.de",
		  "digest":"5d2c6e02a0cced8fb05f32626437e3d23096480b47efbba659b6d9e80c85d280",
		  "severity":"suspend",
		  "comment":"Third-party bots"
		},
		{
		  "domain":"birdbots.leptonics.com",
		  "digest":"ce019d8d32cce8e369ac4367f4dc232103e6f489fbdd247fb99f9c8a646078a4",
		  "severity":"suspend",
		  "comment":"Third-party bots"
		}
	  ]`
)

func TestGetInstanceData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup Instance
		var instance Instance
		err := json.Unmarshal([]byte(testinstance), &instance)
		if err != nil {
			t.Fatalf("error unmarshalling test instance: %v", err)
		}

		// Return based on URI
		switch r.URL.Path {
		case InstanceURI:
			body, err := json.Marshal(instance)
			if err != nil {
				t.Fatalf("error marshalling test instance: %v", err)
			}
			fmt.Fprintln(w, string(body))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Setup client
	client := NewClient(ts.URL)

	_, err := client.GetInstanceData()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}
}

func TestGetInstancePeers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup InstancePeers
		instancepeers := InstancePeers{"tilde.zone", "mspsocial.net", "conf.tube"}

		// Return based on URI
		switch r.URL.Path {
		case InstancePeersURI:
			body, err := json.Marshal(instancepeers)
			if err != nil {
				t.Fatalf("error marshalling test instance peers: %v", err)
			}
			fmt.Fprintln(w, string(body))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Setup client
	client := NewClient(ts.URL)

	ip, err := client.GetInstancePeers()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}

	if len(ip) != 3 {
		t.Fatalf("should have returned 3 peers but instead returned: %d", len(ip))
	}
}

// FIXME - not currently implemented
/*
func TestGetInstancePeersWhitelisted(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup InstancePeers
		instancepeers := InstancePeers{"tilde.zone", "mspsocial.net", "conf.tube"}

		// Setup unauthorized error
		var unauth Unauthorized
		err := json.Unmarshal([]byte(testunauthorized), &unauth)
		if err != nil {
			t.Fatalf("error unmarshalling unauthorized error: %v", err)
		}

		// Return based on URI
		switch r.URL.Path {
		case InstancePeersURI:
			auth := r.Header.Get("Authorization")
			if auth == "" {
				body, err := json.Marshal(unauth)
				if err != nil {
					t.Fatalf("error marshalling unauthorized error: %v", err)
				}
				http.Error(w, string(body), http.StatusUnauthorized)
			}

			body, err := json.Marshal(instancepeers)
			if err != nil {
				t.Fatalf("error marshalling test instance peers: %v", err)
			}
			fmt.Fprintln(w, string(body))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Setup client
	client := NewClient(ts.URL)

	ip, err := client.GetInstancePeers()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}

	if len(ip) != 3 {
		t.Fatalf("should have returned 3 peers but instead returned: %d", len(ip))
	}
}
*/

func TestGetInstanceActivity(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup InstanceActivity
		var instanceactivity InstanceActivity
		err := json.Unmarshal([]byte(testinstanceactivity), &instanceactivity)
		if err != nil {
			t.Fatalf("error unmarshalling test instance activity: %v", err)
		}

		// Return based on URI
		switch r.URL.Path {
		case InstanceActivityURI:
			body, err := json.Marshal(instanceactivity)
			if err != nil {
				t.Fatalf("error marshalling test instance activity: %v", err)
			}
			fmt.Fprintln(w, string(body))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Setup client
	client := NewClient(ts.URL)

	ia, err := client.GetInstanceActivity()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}

	if len(ia) != 12 {
		t.Fatalf("should have returned 12 weeks but instead returned: %d", len(ia))
	}
}

// FIXME - not currently implemented
/*
func TestGetInstanceActivityWhitelisted(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup InstanceActivity
		var instanceactivity InstanceActivity
		err := json.Unmarshal([]byte(testinstanceactivity), &instanceactivity)
		if err != nil {
			t.Fatalf("error unmarshalling test instance activity: %v", err)
		}

		// Setup unauthorized error
		var unauth Unauthorized
		err = json.Unmarshal([]byte(testunauthorized), &unauth)
		if err != nil {
			t.Fatalf("error unmarshalling unauthorized error: %v", err)
		}

		// Return based on URI
		switch r.URL.Path {
		case InstanceActivityURI:
			auth := r.Header.Get("Authorization")
			if auth == "" {
				body, err := json.Marshal(unauth)
				if err != nil {
					t.Fatalf("error marshalling unauthorized error: %v", err)
				}
				http.Error(w, string(body), http.StatusUnauthorized)
			}

			body, err := json.Marshal(instanceactivity)
			if err != nil {
				t.Fatalf("error marshalling test instance activity: %v", err)
			}
			fmt.Fprintln(w, string(body))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Setup client
	client := NewClient(ts.URL)

	ia, err := client.GetInstanceActivity()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}

	if len(ia) != 12 {
		t.Fatalf("should have returned 12 weeks but instead returned: %d", len(ia))
	}
}
*/

func TestGetInstanceRules(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup InstanceRules
		var instancerules InstanceRules
		err := json.Unmarshal([]byte(testinstancerules), &instancerules)
		if err != nil {
			t.Fatalf("error unmarshalling test instance rules: %v", err)
		}

		// Return based on URI
		switch r.URL.Path {
		case InstanceRulesURI:
			body, err := json.Marshal(instancerules)
			if err != nil {
				t.Fatalf("error marshalling test instance rules: %v", err)
			}
			fmt.Fprintln(w, string(body))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Setup client
	client := NewClient(ts.URL)

	ir, err := client.GetInstanceRules()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}

	if len(ir) != 6 {
		t.Fatalf("should have returned 6 rules but instead returned: %d", len(ir))
	}
}

func TestGetInstanceDomainsBlocked(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup InstanceDomainsBlocked
		var blockeddomains DomainsBlocked
		err := json.Unmarshal([]byte(testinstancedomainsblocked), &blockeddomains)
		if err != nil {
			t.Fatalf("error unmarshalling test instance domains blocked: %v", err)
		}

		// Return based on URI
		switch r.URL.Path {
		case InstanceDomainsBlockedyURI:
			body, err := json.Marshal(blockeddomains)
			if err != nil {
				t.Fatalf("error marshalling test instance domains blocked: %v", err)
			}
			fmt.Fprintln(w, string(body))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Setup client
	client := NewClient(ts.URL)

	db, err := client.GetInstanceDomainsBlocked()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}

	if len(db) != 2 {
		t.Fatalf("should have returned 2 blocked domains but instead returned: %d", len(db))
	}
}

// FIXME - not currently implemented
/*
func TestGetInstanceDomainsBlockedWhitelisted(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup InstanceDomainsBlocked
		var blockeddomains DomainsBlocked
		err := json.Unmarshal([]byte(testinstancedomainsblocked), &blockeddomains)
		if err != nil {
			t.Fatalf("error unmarshalling test instance domains blocked: %v", err)
		}

		// Setup unauthorized error
		var unauth Unauthorized
		err = json.Unmarshal([]byte(testunauthorized), &unauth)
		if err != nil {
			t.Fatalf("error unmarshalling unauthorized error: %v", err)
		}

		// Return based on URI
		switch r.URL.Path {
		case InstanceDomainsBlockedyURI:
			auth := r.Header.Get("Authorization")
			if auth == "" {
				body, err := json.Marshal(unauth)
				if err != nil {
					t.Fatalf("error marshalling unauthorized error: %v", err)
				}
				http.Error(w, string(body), http.StatusUnauthorized)
			}

			body, err := json.Marshal(blockeddomains)
			if err != nil {
				t.Fatalf("error marshalling test instance domains blocked: %v", err)
			}
			fmt.Fprintln(w, string(body))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Setup client
	client := NewClient(ts.URL)

	db, err := client.GetInstanceDomainsBlocked()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}

	if len(db) != 2 {
		t.Fatalf("should have returned 2 blocked domains but instead returned: %d", len(db))
	}
}
*/
