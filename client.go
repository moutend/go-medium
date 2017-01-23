// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type rawbody []byte

func (r rawbody) User(c *Client) (*User, error) {
	var i struct {
		Data User
	}
	err := decodeJSON(bytes.NewReader(r), &i)
	i.Data.client = c
	return &i.Data, err
}

// Client represents Medium API client.
type Client struct {
	Root              *url.URL
	ApplicationID     string
	ApplicationSecret string
	AccessToken       string

	httpClient *http.Client
	logger     *log.Logger
	name       string
	version    string
}

// NewClient returns API client.
func NewClient(logger *log.Logger) (c *Client) {
	if logger == nil {
		logger = log.New(ioutil.Discard, "discard logging messages", log.LstdFlags)
	}
	u, _ := url.Parse("https://api.medium.com/v1")
	token := os.Getenv("MEDIUM_API_TOKEN")
	return &Client{
		Root:        u,
		AccessToken: token,
		httpClient:  &http.Client{},
		logger:      logger,
		name:        "go-medium",
		version:     version,
	}

}

func (c *Client) newRequest(method string, u *url.URL, body io.Reader) (req *http.Request, err error) {
	if u.Host == "" {
		u, err = url.Parse(c.Root.String() + u.String())
		if err != nil {
			return
		}
	}

	c.logger.Printf("%s %s\n", method, u.String())
	req, err = http.NewRequest(method, u.String(), body)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", c.name+"/"+c.version)
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")
	return
}

func (c *Client) do(req *http.Request) (res *http.Response, err error) {
	if res, err = c.httpClient.Do(req); err != nil {
		return
	}
	c.logger.Printf("%s %s\n", res.Proto, res.Status)
	return
}

func (c *Client) get(u *url.URL) (r rawbody, err error) {
	req, err := c.newRequest("GET", u, nil)
	if err != nil {
		return
	}

	res, err := c.do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err = ioutil.ReadAll(res.Body)
	return
}
func (c *Client) post(u *url.URL, post Post) (r rawbody, err error) {
	body, err := json.Marshal(post)
	if err != nil {
		return
	}
	req, err := c.newRequest("POST", u, bytes.NewReader(body))
	if err != nil {
		return
	}

	res, err := c.do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err = ioutil.ReadAll(res.Body)
	return
}

// User returns the authenticated user's details.
func (c *Client) User() (u *User, err error) {
	path, _ := url.Parse("/me")
	r, err := c.get(path)
	if err != nil {
		return
	}
	return r.User(c)
	return
}
