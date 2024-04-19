package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) ResolveAdaHandle(handle string) (*models.BasicResponse, error) {
	url := fmt.Sprintf("/ecosystem/adahandle/%s", handle)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close()
	var adaHandle models.BasicResponse
	err = json.NewDecoder(resp.Body).Decode(&adaHandle)
	if err != nil {
		return nil, err
	}
	return &adaHandle, nil
}
