package client

import (
	"encoding/json"
	"fmt"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) TxManagerHistory() (*[]models.TxManagerState, error) {
	url := "/txmanager/history"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var txManagerStates []models.TxManagerState
	err = json.NewDecoder(resp.Body).Decode(&txManagerStates)
	if err != nil {
		return nil, err
	}
	return &txManagerStates, nil
}

func (c *Client) TxManagerSubmit(cbor string) (models.BasicResponse, error) {
	url := "/txmanager"
	resp, err := c.post(url, cbor)
	if err != nil {
		return models.BasicResponse{}, err
	}
	defer resp.Body.Close()
	var submitTx models.BasicResponse
	err = json.NewDecoder(resp.Body).Decode(&submitTx)
	if err != nil {
		return models.BasicResponse{}, err
	}
	return submitTx, nil
}

func (c *Client) TxManagerSubmitTurbo(cbor string) (models.BasicResponse, error) {
	url := "/txmanager/turbosubmit"
	resp, err := c.post(url, cbor)
	if err != nil {
		return models.BasicResponse{}, err
	}
	defer resp.Body.Close()
	var submitTx models.BasicResponse
	err = json.NewDecoder(resp.Body).Decode(&submitTx)
	if err != nil {
		return models.BasicResponse{}, err
	}
	return submitTx, nil
}

func (c *Client) TxManagerState(txHash string) (*models.TxManagerState, error) {
	url := fmt.Sprintf("/txmanager/%s/state", txHash)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var txManagerState models.TxManagerState
	err = json.NewDecoder(resp.Body).Decode(&txManagerState)
	if err != nil {
		return nil, err
	}
	return &txManagerState, nil
}
