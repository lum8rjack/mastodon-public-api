package mastodon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCustomEmojis(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup Emojis
		emojis := Emojis{
			{
				Shortcode:       "blobaww",
				URL:             "https://files.mastodon.social/custom_emojis/images/000/011/739/original/blobaww.png",
				StaticURL:       "https://files.mastodon.social/custom_emojis/images/000/011/739/static/blobaww.png",
				VisibleInPicker: true,
				Category:        "Blobs",
			},
			{
				Shortcode:       "aaaa",
				URL:             "https://files.mastodon.social/custom_emojis/images/000/007/118/original/aaaa.png",
				StaticURL:       "https://files.mastodon.social/custom_emojis/images/000/007/118/static/aaaa.png",
				VisibleInPicker: true,
			},
		}

		// Return based on URI
		switch r.URL.Path {
		case CustomEmojisURI:
			body, err := json.Marshal(emojis)
			if err != nil {
				t.Fatalf("error marshalling emoji: %v", err)
			}
			fmt.Fprintln(w, string(body))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Setup client
	client, err := NewClient(ts.URL)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	emojis, err := client.GetCustomEmojis()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}

	if len(emojis) != 2 {
		t.Fatalf("should return 2 emojis, instead go %d", len(emojis))
	}
}
