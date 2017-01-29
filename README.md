# go-medium

[![GitHub release](https://img.shields.io/github/release/moutend/go-medium.svg?style=flat-square)][release]
[![Go Documentation](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![CircleCI](https://circleci.com/gh/moutend/go-medium.svg?style=svg&circle-token=2227e983ccd1eafc5f05f66ade52b7030ff0b18c)](https://circleci.com/gh/moutend/go-medium)

[release]: https://github.com/moutend/go-medium/releases
[godocs]: https://godoc.org/github.com/moutendgo-medium/
[license]: https://github.com/moutend/go-medium/blob/master/LICENSE
[status]: https://circleci.com/gh/moutend/go-medium

`go-medium` is Medium API client for Go.

As you know, Medium provides official API client for Go, literally named [medium-sdk-go](https://github.com/Medium/medium-sdk-go).
However, that client seems to be not able to publish an article to the authorized user's publications, I decided to implement alternative Go client for Medium.

# Installation

```shell
$ go install github.com/moutend/go-medium
```

# Usage

In the following example, an article written in markdown will be published at the publication owned by authorized user.

```go
package main

import (
	"fmt"
	"log"

	medium "github.com/moutend/go-medium"
)

func main() {
	clientID := "XXXXXXXX"
	clientSecret := "YYYYYYYY"
	accessToken := "ZZZZZZZ"
	c := medium.NewClient(clientID, clientSecret, accessToken)
	u, err := c.User()
	if err != nil {
		log.Fatal(err)
	}
	// Get detail of all publications.
	ps, err := u.Publications()
	if err != nil {
		log.Fatal(err)
	}
	// If the user has no publications yet, exit with error message.
	if len(ps) == 0 {
		log.Fatalf("%s has no publications yet.\n", u.Username)
	}
	// Post an article to the first publication.
	a, err := ps[0].Post(medium.Article{
		Title:         "test",
		ContentFormat: "markdown",
		Content: `# Test article

This article was posted from command line.`,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published an article at %s\n", a.URL)
}
```

If you are using self-issued token, please specify empty string `""` as the client ID and client secret.

# LICENSE

MIT
