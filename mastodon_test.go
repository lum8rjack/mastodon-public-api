package mastodon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	server := "https://infosec.exchange"
	client := NewClient(server)

	if client.Server != server {
		t.Fatalf("server was incorrectly set to: %s", client.Server)
	}

	if client.Timeout != time.Second*Timeout {
		t.Fatalf("timeout was incorrectly set to: %s", client.Timeout)
	}

	if client.UserAgent != UserAgent {
		t.Fatalf("user-agent was incorrectly set to: %s", client.UserAgent)
	}
}

func TestSendRequest(t *testing.T) {
	// Setup return status
	type Status struct {
		Status string
	}

	returnStatus := Status{Status: "success"}
	rs, err := json.Marshal(returnStatus)
	if err != nil {
		t.Fatalf("error setting up status: %v", err)
	}

	// Setup http server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Return based on URI
		switch r.URL.Path {
		case "/":
			fmt.Fprintln(w, string(rs))
			return
		}

		// URI not specified above, return status not found
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))
	defer ts.Close()

	// Create client
	c := NewClient(ts.URL)

	// Check base url
	body, err := c.SendRequest(ts.URL)
	if err != nil {
		t.Fatalf("should not fail: %v", err)
	}
	var rstatus Status
	err = json.Unmarshal(body, &rstatus)
	if err != nil {
		t.Fatalf("invalid return status: %v", err)
	}
	if rstatus.Status != "success" {
		t.Fatalf("should be success instead got: %s", rstatus.Status)
	}

}
