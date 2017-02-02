# go-medium

[![GitHub release](https://img.shields.io/github/release/moutend/go-medium.svg?style=flat-square)][release]
[![Go Documentation](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![CircleCI](https://circleci.com/gh/moutend/go-medium.svg?style=svg&circle-token=2227e983ccd1eafc5f05f66ade52b7030ff0b18c)](https://circleci.com/gh/moutend/go-medium)

[release]: https://github.com/moutend/go-medium/releases
[godocs]: https://godoc.org/github.com/moutendgo-medium/
[license]: https://github.com/moutend/go-medium/blob/master/LICENSE
[status]: https://circleci.com/gh/moutend/go-medium

`go-medium` is a Medium API client for Go.

As you know, Medium provides an official API client for Go, literally named [medium-sdk-go](https://github.com/Medium/medium-sdk-go).
However, that client is not able to publish an article to the authorized user's publications.
I decided to implement alternative Go client for Medium.

# Installation

```shell
$ go install github.com/moutend/go-medium
```

# Usage

## Initialize client

First off, you need initialize a client like this:

```go
c := medium.NewClient(clientID, clientSecret, accessToken)
```

It's not recommended  but if you want to use self-issued token, set `clientID` and `clientSecret` blank.

## Example

In the following example, it demonstrates the steps below:

- Initialize a client.
- Get information about current authorized user.
- Get publications owned by current authorized user.
- Create an article as Markdown format.
- Post the article to the first publication.

```go
package main

import (
	"fmt"
	"log"

	medium "github.com/moutend/go-medium"
)

func main() {
	// It's not recommended, but it uses self-issued token in this example.
	accessToken := "self-issued-access-token"
	c := medium.NewClient("", "", accessToken)

	// Get information about current authorized user.
	// See user.go to check all exported fields.
	u, err := c.User()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You are logged in as %s.", u.Username)

	// Get list of publications.
	ps, err := u.Publications()
	if err != nil {
		log.Fatal(err)
	}
	// It assumes that the user has one or more publications.
	// See publication.go to check all exported fields.
	if len(ps) == 0 {
		log.Fatalf("%s has no publications yet.\n", u.Username)
	}
	fmt.Printf("You have a publication named %s.\n", ps[0].Name)

	// Publish an article as Markdown format.
	// For more dail, see next section.
	article := medium.Article{
		Title:         "test",
		ContentFormat: "markdown",
		Content: `# Test

## Sub title

# Using Medium API

This article was published from command line.`,
	}
	// Publish an article to the first publication.
	// See article.go to check all exported fields of PostedArticle.
	pa, err := ps[0].Post(article)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published an article at %s\n", a.URL)
}
```

# Notes

## Published date and time of Article

If you want to specify the field `PublishedAt` of Article, it must be formatted according to the layout below:

```
2006-01-02T15:04:05+07000
```

Note that you cannot specify the date and time after UTC+07:00.
I don't know why but it seems to be specification of the Medium API v1.
For example, Japan standard Time is UTC+09:00, the article posted from machine which timezone is set as JST will be rejected.
Also, you cannot specify the date and time before Jan 1st, 1970.

The `PublishedAt` field will be set as current time if you didn't specify it.

## Setting up API token with OAuth

`go-medium` doesn't provide the method for obtaining API token with OAuth.
However, it has `Token` method which receives shortlive code and redirect URI, and then generate a new token for Medium API.

One way to generate API token with OAuth is launching the local web server and redirect HTTP request to that web server.

# LICENSE

MIT
