package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	username     string
	password     string
	organization string
	project      string
	apiVersion   string
	httpClient   *http.Client
}

func NewClient(username string, password string, organization string, project string) *Client {
	return &Client{
		username:     username,
		password:     password,
		organization: organization,
		project:      project,
		apiVersion:   "6.0-preview.1",
		httpClient:   &http.Client{},
	}
}

func (c *Client) HttpRequest(path string, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/%s?api-version=%s", c.organization, c.project, path, c.apiVersion)
}
