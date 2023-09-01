package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/maestro-org/go-sdk/config"
)

type Client struct {
	apiKey     string
	network    string
	version    string
	HTTPClient *http.Client
	baseUrl    string
}

func NewClient(apiKey string, network string) *Client {
	cfg := config.GetConfig()
	paths := []string{"https:///", fmt.Sprintf("%s.gomaestro-api.org", strings.ToLower(network)), cfg.Client.Version}
	return &Client{
		apiKey:  apiKey,
		network: network,
		version: cfg.Client.Version,
		HTTPClient: &http.Client{
			Timeout: time.Duration(cfg.Client.Timeout),
		},
		baseUrl: path.Join(paths...),
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Unmarshall and populate v
	fullResponse := successResponse{
		Data: v,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}
