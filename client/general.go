package client

import (
	"encoding/json"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) ChainTip() (*models.ChainTip, error) {
	url := "/chain-tip"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var chainTip models.ChainTip
	err = json.NewDecoder(resp.Body).Decode(&chainTip)
	if err != nil {
		return nil, err
	}
	return &chainTip, nil
}

func (c *Client) EraHistory() (*models.EraHistory, error) {
	url := "/era-history"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var eraHistory models.EraHistory
	err = json.NewDecoder(resp.Body).Decode(&eraHistory)
	if err != nil {
		return nil, err
	}
	return &eraHistory, nil
}

func (c *Client) ProtocolParameters() (*models.ProtocolParameters, error) {
	url := "/protocol-params"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var protocolParameters models.ProtocolParameters
	err = json.NewDecoder(resp.Body).Decode(&protocolParameters)
	if err != nil {
		return nil, err
	}
	return &protocolParameters, nil
}

func (c *Client) BlockChainStartTime() (models.BasicResponse, error) {
	url := "/system-start"
	resp, err := c.get(url)
	if err != nil {
		return models.BasicResponse{}, err
	}
	defer resp.Body.Close()
	var blockchainStartTime models.BasicResponse
	err = json.NewDecoder(resp.Body).Decode(&blockchainStartTime)
	if err != nil {
		return models.BasicResponse{}, err
	}
	return blockchainStartTime, nil
}
