package mastodon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testtrendslinks string = `[
		{
		  "url": "https://www.nbcnews.com/specials/plan-your-vote-2022-elections/index.html",
		  "title": "Plan Your Vote: 2022 Elections",
		  "description": "Everything you need to know about the voting rules where you live, including registration, mail-in voting, changes since 2020, and more.",
		  "type": "link",
		  "author_name": "NBC News",
		  "author_url": "",
		  "provider_name": "NBC News",
		  "provider_url": "",
		  "html": "",
		  "width": 400,
		  "height": 225,
		  "image": "https://files.mastodon.social/cache/preview_cards/images/045/027/478/original/0783d5e91a14fd49.jpeg",
		  "embed_url": "",
		  "blurhash": "UcQmF#ay~qofj[WBj[j[~qof9Fayofofayay",
		  "history": [
			{
			  "day": "1661817600",
			  "accounts": "7",
			  "uses": "7"
			},
			{
			  "day": "1661731200",
			  "accounts": "23",
			  "uses": "23"
			},
			{
			  "day": "1661644800",
			  "accounts": "0",
			  "uses": "0"
			},
			{
			  "day": "1661558400",
			  "accounts": "0",
			  "uses": "0"
			},
			{
			  "day": "1661472000",
			  "accounts": "0",
			  "uses": "0"
			},
			{
			  "day": "1661385600",
			  "accounts": "0",
			  "uses": "0"
			},
			{
			  "day": "1661299200",
			  "accounts": "0",
			  "uses": "0"
			}
		  ]
		}
	  ]`

	testtrendstags string = `[
		{
		  "name": "hola",
		  "url": "https://mastodon.social/tags/hola",
		  "history": [
			{
			  "day": "1574726400",
			  "uses": "13",
			  "accounts": "10"
			}
		  ]
		},
		{
		  "name": "SaveDotOrg",
		  "url": "https://mastodon.social/tags/SaveDotOrg",
		  "history": [
			{
			  "day": "1574726400",
			  "uses": "9",
			  "accounts": "9"
			}
		  ]
		},
		{
		  "name": "introduction",
		  "url": "https://mastodon.social/tags/introduction",
		  "history": [
			{
			  "day": "1574726400",
			  "uses": "15",
			  "accounts": "14"
			}
		  ]
		}
	  ]`
)

func TestGetTrendsLinks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup TrendsLinks
		var trendslinks TrendLinks
		err := json.Unmarshal([]byte(testtrendslinks), &trendslinks)
		if err != nil {
			t.Fatalf("error unmarshalling test trends links: %v", err)
		}

		// Return based on URI
		switch r.URL.Path {
		case TrendsLinksURI:
			body, err := json.Marshal(trendslinks)
			if err != nil {
				t.Fatalf("error marshalling trends links: %v", err)
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

	_, err = client.GetTrendsLinks()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}
}

func TestGetTrendsTags(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setup TrendsTags
		var trendstags TrendTags
		err := json.Unmarshal([]byte(testtrendstags), &trendstags)
		if err != nil {
			t.Fatalf("error unmarshalling test trends tags: %v", err)
		}

		// Return based on URI
		switch r.URL.Path {
		case TrendsTagsURI:
			body, err := json.Marshal(trendstags)
			if err != nil {
				t.Fatalf("error marshalling trends tags: %v", err)
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

	_, err = client.GetTrendsTags()
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}
}
