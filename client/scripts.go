package client

import (
	"encoding/json"
	"fmt"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) ScriptByHash(hash string) (*models.ScriptByHash, error) {
	url := fmt.Sprintf("/scripts/%s", hash)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var scriptByHash models.ScriptByHash
	err = json.NewDecoder(resp.Body).Decode(&scriptByHash)
	if err != nil {
		return nil, err
	}
	return &scriptByHash, nil

}
