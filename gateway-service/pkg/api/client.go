package api

import (
	"fmt"
	"net/http"
)

type BaseClient struct {
	HTTPClient *http.Client
}

func (c *BaseClient) SendRequest(req *http.Request) (resp *http.Response, err error) {
	if c.HTTPClient == nil {
		return resp, fmt.Errorf("no http client")
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err = c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return resp, fmt.Errorf("failed to send request. error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("status code !ok")
	}

	return resp, err
}
