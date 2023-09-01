package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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
	return &Client{
		apiKey:  apiKey,
		network: network,
		version: cfg.Client.Version,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		baseUrl: fmt.Sprintf("https://%s.gomaestro-api.org/%s", network, cfg.Client.Version),
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// type successResponse struct {
// 	Code int         `json:"code"`
// 	Data interface{} `json:"data"`
// }

func (c *Client) sendRequest(req *http.Request, responseBody *string) error {
	req.Header.Set("Accept", "application/json")
	req.Header.Add("api-key", c.apiKey)

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

	respBodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Errorf("failed to read body: %s", err)
	}
	defer res.Body.Close()

	*responseBody = string(respBodyBytes)

	return nil
}
