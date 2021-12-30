package twentynine

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	url  string
	http http.Client
}

type ClientOption func(cli *Client)

func OptTimeout(timeout time.Duration) ClientOption {
	return func(cli *Client) { cli.http.Timeout = timeout }
}

func OptURL(rawurl string) ClientOption {
	return func(cli *Client) { cli.url = rawurl }
}

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		url: TwentyNineApiUrl,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (cli *Client) Links() (links []Link, err error) {
	r, err := cli.http.Get(cli.url)
	if err != nil {
		return
	}
	if r.StatusCode != http.StatusOK {
		msgBuf, _ := ioutil.ReadAll(r.Body)
		err = Error{
			Code:    r.StatusCode,
			Message: string(msgBuf),
		}
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&links); err != io.EOF {
		return
	}
	err = nil // err ?= io.EOF
	return
}

func (cli *Client) LinkLink(link Link) (err error) {
	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(link); err != nil {
		return
	}
	r, err := cli.http.Post(cli.url, "application/json", buf)
	if err != nil {
		return
	}
	if r.StatusCode != http.StatusOK {
		msgBuf, _ := ioutil.ReadAll(r.Body)
		err = Error{
			Code:    r.StatusCode,
			Message: string(msgBuf),
		}
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&link); err != io.EOF {
		return
	}
	err = nil // err ?= io.EOF
	return
}
