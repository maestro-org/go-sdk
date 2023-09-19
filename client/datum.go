package client

import (
	"encoding/json"
	"fmt"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) DatumFromHash(hash string) (*models.DatumFromHash, error) {
	url := fmt.Sprintf("/data/%s", hash)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var datum models.DatumFromHash
	err = json.NewDecoder(resp.Body).Decode(&datum)
	if err != nil {
		return nil, err
	}
	return &datum, nil
}
