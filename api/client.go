package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Endpoint string
	Username string
	Password string
}

func (c *Client) PostRequest(url string, i interface{}) (*http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(i)
	return c.Request("POST", url, 201, b)
}

func (c *Client) GetRequest(url string) (*http.Response, error) {
	return c.Request("GET", url, 200, nil)
}

func (c *Client) Request(verb, url string, code int, payload io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(verb, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != code {
		return resp, errors.New(fmt.Sprintf("Received %d, expecting %d status code while fetching %s", resp.StatusCode, code, url))
	}
	return resp, err
}
