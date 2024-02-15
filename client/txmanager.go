package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

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

func (c *Client) TxManagerSubmit(txHex string) (string, error) {
	url := "/txmanager"
	txBuffer, _ := hex.DecodeString(txHex)
	resp, err := c.postBuffer(url, txBuffer)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	txHash := string(bodyBytes)
	return txHash, nil
}

func (c *Client) TxManagerSubmitTurbo(txHex string) (string, error) {
	url := "/txmanager/turbosubmit"
	txBuffer, _ := hex.DecodeString(txHex)
	resp, err := c.postBuffer(url, txBuffer)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	txHash := string(bodyBytes)
	return txHash, nil
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
