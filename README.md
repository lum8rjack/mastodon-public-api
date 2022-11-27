# mastodon-public-api
API to fetch public data from Mastodon servers.

Mastodon documentation: https://docs.joinmastodon.org/client/public/
Additional Go Libraries: https://docs.joinmastodon.org/client/libraries/#go

## Usage

Example program to get instance data.

```go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lum8rjack/mastodon-public-api"
)

func main() {
	server := "https://infosec.exchange"

	client := mastodon.NewClient(*server)

	instanceData, err := client.GetInstanceData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Version: %s\n", instanceData.Version)
	fmt.Printf("User count: %d\n", instanceData.Stats.UserCount)
	fmt.Printf("Description: %s\n", instanceData.Description)
}
```

Output:
```bash
Version: 4.0.2+glitch
User count: 32487
Description: This is a Mastodon instance open to the general public, but may contain more than the usual amount of IT security discussions.
```

## Status of implementations

* [ ] GET /api/v1/accounts/:id
* [ ] GET /api/v1/accounts/:id/statuses
* [x] GET /api/v1/custom_emojis
* [ ] GET /api/v1/directory
* [x] GET /api/v1/instance
* [x] GET /api/v1/instance/activity
* [x] GET /api/v1/instance/domain_block
* [x] GET /api/v1/instance/peers
* [x] GET /api/v1/instance/rules
* [ ] GET /api/v1/polls/:id
* [ ] GET /api/v1/statuses/:id
* [ ] GET /api/v1/statuses/:id/context
* [ ] GET /api/v1/statuses/:id/favourited_by
* [ ] GET /api/v1/statuses/:id/reblogged_by
* [ ] GET /api/v1/timelines/public
* [ ] GET /api/v1/timelines/tag/:hashtag
* [x] GET /api/v1/trends/links
* [ ] GET /api/v1/trends/statuses
* [x] GET /api/v1/trends/tags


## License

MIT
