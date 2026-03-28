package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/unicode/norm"
)

const (
	// TODO: Replace with your API base URL
	baseURL   = "https://api.example.com/search"
	timeout   = 10 * time.Second
	userAgent = "{{PROJECT_NAME}}-cli"
)

// Result is the data returned by the API.
// TODO: Replace with your actual API response structure.
type Result struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Response is the API response wrapper.
type Response struct {
	Error   string   `json:"error,omitempty"`
	Count   int      `json:"count"`
	Results []Result `json:"results"`
}

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{Timeout: timeout},
	}
}

func (c *Client) Search(ctx context.Context, keyword string) ([]Result, error) {
	normalized := norm.NFC.String(keyword)
	u := fmt.Sprintf("%s?q=%s", baseURL, url.QueryEscape(normalized))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{
			Message:    string(body),
			StatusCode: resp.StatusCode,
		}
	}

	var apiResp Response
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	if apiResp.Error != "" {
		return nil, &APIError{Message: apiResp.Error}
	}

	return apiResp.Results, nil
}
