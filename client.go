// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

// Package medium provides thin wrapper for Medium API.
package medium

import (
	"bytes"
	"encoding/json"
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
	"time"
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

func (c *Client) postArticle(mode, id string, a Article) (*PostedArticle, error) {
	if a.PublishedAt == "" {
		// I don't know why but Medium API doesn't accept the article published after UTC+07:00.
		// For example, Japan Standard Time is UTC+09:00, so that the article posted from Japan will be rejected.
		// We need set current UTC timezone to Etc/GMT-7 (+07:00) before formatting date and time.
		local, err := time.LoadLocation("Etc/GMT-7")
		if err != nil {
			return nil, err
		}
		a.PublishedAt = time.Now().In(local).Format("2006-01-02T15:04:05+07:00")
	}
	if a.PublishStatus == "" {
		a.PublishStatus = "public"
	}
	if a.Title == "" {
		a.Title = "Untitled"
	}
	if a.ContentFormat != "html" && a.ContentFormat != "markdown" {
		return nil, fmt.Errorf("only \"html\" and \"markdown\" are valid as content format")
	}
	if a.Content == "" {
		return nil, fmt.Errorf("content should not be blank")
	}
	c.logger.Println(a.PublishStatus, a.PublishedAt)
	content, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/%s/%s/posts", mode, id)
	req, err := c.newRequest("POST", path, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	r, err := c.do(req)
	if err != nil {
		return nil, err
	}
	return r.PostedArticle(c)
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
