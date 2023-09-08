package client

import (
	"fmt"
	"net/http"
)

type HttpClient struct {
	BaseURL       string
	GetEndpoints  []string
	PostEndpoints []string
	HeadEndpoints []string
	Client        *http.Client
}

func NewHTTPClient(baseURL string, getEndpoints, postEndpoints, headEndpoints []string) *HttpClient {
	return &HttpClient{
		GetEndpoints:  getEndpoints,
		PostEndpoints: postEndpoints,
		HeadEndpoints: headEndpoints,
		Client:        &http.Client{},
	}
}

func (c *HttpClient) SendRequest(req *http.Request) error {
	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s to %v request failed with status code: %d", req.Method, req.URL, res.StatusCode)
	}
	return nil
}

func (c *HttpClient) PostRequests() error {

	return nil
}
