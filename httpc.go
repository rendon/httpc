// Package httpc provides an HTTP client interface for creating API clients. It
// comes with a default implementation that covers most cases.
//
// Using an interface for your API client allows you to stub requests in tests,
// just create a type that implements the interface with your custom behavior
// and the rest of your application will work seamlessly.
//
// This package aims to be minimal, just a bit more than the standard HTTP
// package so you can still be in control of your software.
package httpc

import (
	"bytes"
	"net/http"
	"time"
)

// WebAPIClient is an interface for Web API clients.
type WebAPIClient interface {
	Get(url string) (*http.Response, error)
	Post(url string, body []byte) (*http.Response, error)
	Put(url string, body []byte) (*http.Response, error)
	Delete(url string) (*http.Response, error)
}

// Client type implements WebAPIClient with a basic behaviour, assuming
// "application/json" as Content type.
type Client struct {
	httpClient http.Client
}

// Get performs an HTTP GET request, "application/json" is assumed.
func (t Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	return t.httpClient.Do(req)
}

// Get package function for DefaultClient.Get
func Get(url string) (*http.Response, error) {
	return DefaultClient.Get(url)
}

// Post performs an HTTP POST request, "application/json" is assumed.
func (t Client) Post(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	return t.httpClient.Do(req)
}

// Post package function for DefaultClient.Post
func Post(url string, body []byte) (*http.Response, error) {
	return DefaultClient.Post(url, body)
}

// Put performs an HTTP PUT request, "application/json" is assumed.
func (t Client) Put(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	return t.httpClient.Do(req)
}

// Put package function for DefaultClient.Put
func Put(url string, body []byte) (*http.Response, error) {
	return DefaultClient.Put(url, body)
}

// Delete performs an HTTP DELETE request, "application/json" is assumed.
func (t Client) Delete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	return t.httpClient.Do(req)
}

// Delete package function for DefaultClient.Delete
func Delete(url string) (*http.Response, error) {
	return DefaultClient.Delete(url)
}

// NewClient returns an initialized client with the given timeout, a timeout of
// zero means no timeout, see https://golang.org/pkg/net/http/#Client.
func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: http.Client{Timeout: timeout},
	}
}

// DefaultClient a Client type with a default implementation for WebAPIClient.
var DefaultClient = &Client{}
