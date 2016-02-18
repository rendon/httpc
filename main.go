package httpc

import (
	"bytes"
	"net/http"
)

type WebAPIClient interface {
	Get(url string) (*http.Response, error)
	Post(url string, body []byte) (*http.Response, error)
}

type Client struct{}

func (t Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func (t Client) Post(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

var DefaultClient = &Client{}
