
package engine

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	URL        string
	StatusCode int
	Status     string
	Headers    http.Header
	Duration   time.Duration
}

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *HTTPClient) Get(ctx context.Context, target string) (*Response, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		target,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	start := time.Now()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	return &Response{
		URL:        target,
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    resp.Header.Clone(),
		Duration:   time.Since(start),
	}, nil
}

