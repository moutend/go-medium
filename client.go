// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

// Package medium provides thin wrapper for Medium API.
package medium

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"strings"
)

// Error represents medium API.
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
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
func NewClient(clientID, clientSecret, accessToken string) (c *Client) {
	u, _ := url.Parse("https://api.medium.com/v1")
	return &Client{
		Root:              u,
		AccessToken:       accessToken,
		ApplicationID:     clientID,
		ApplicationSecret: clientSecret,
		httpClient:        &http.Client{},
		logger:            log.New(ioutil.Discard, "discard logging messages", log.LstdFlags),
		name:              "go-medium",
		version:           version,
	}
}

func (c *Client) newRequest(method, path string, body io.Reader) (req *http.Request, err error) {
	u, err := url.Parse(path)
	if err != nil {
		return
	}
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

func (c *Client) do(req *http.Request) (r rawbody, err error) {
	for k, v := range req.Header {
		c.logger.Printf("%s: %s", k, strings.Join(v, ""))
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	c.logger.Println("@@@")
	c.logger.Printf("%s %s\n", res.Proto, res.Status)
	if r, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}
	if res.StatusCode >= 400 {
		err = r.Error()
		return
	}
	return
}

// User returns the authenticated user's details.
func (c *Client) User() (u *User, err error) {
	req, err := c.newRequest("GET", "/me", nil)
	if err != nil {
		return
	}
	r, err := c.do(req)
	if err != nil {
		return
	}
	return r.User(c)
}

func readImageFile(filename string) (body []byte, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	buf := bytes.Buffer{}
	w := multipart.NewWriter(&buf)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="image"; filename="%s"`, filename))
	h.Set("Content-Type", "image/jpeg")

	part, err := w.CreatePart(h)
	if err != nil {
		return
	}
	if _, err = io.Copy(part, f); err != nil {
		return
	}
	w.Close()
	return buf.Bytes(), nil
}

// Image upload an image.
func (c *Client) Image(filename string) (i *Image, err error) {
	body, err := readImageFile(filename)
	if err != nil {
		return
	}
	req, err := c.newRequest("POST", "/images", bytes.NewReader(body))
	if err != nil {
		return
	}
	r, err := c.do(req)
	if err != nil {
		return
	}
	return r.Image()
}

// SetLogger sets logger.
func (c *Client) SetLogger(logger *log.Logger) {
	c.logger = logger
	return
}
func (c *Client) Token(code, redirectURI string) (token *Token, err error) {
	body := strings.NewReader(fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s", code, c.ApplicationID, c.ApplicationSecret, url.QueryEscape(redirectURI)))
	req, err := c.newRequest("POST", "/tokens", body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Del("Authorization")
	r, err := c.do(req)
	if err != nil {
		return
	}
	return r.Token()
}
