package client

import (
	"encoding/json"
	"fmt"

	"github.com/maestro-org/go-sdk/models"
)

type CreateLock struct {
	Sender                   string `json:"sender"`
	Beneficiary              string `json:"beneficiary"`
	AssetPolicyId            string `json:"asset_policy_id"`
	AssetTokenName           string `json:"asset_token_name"`
	TotalVestingQuantity     int64  `json:"total_vesting_quantity"`
	VestingPeriodStart       int64  `json:"vesting_period_start"`
	VestingPeriodEnd         int64  `json:"vesting_period_end"`
	FirstUnlockPossibleAfter int64  `json:"first_unlock_possible_after"`
	TotalInstallments        int64  `json:"total_installments"`
}

func (c *Client) LockAssets(createLock CreateLock) (*models.LockTransaction, error) {
	url := "/contracts/vesting/lock"

	body := createLock
	resp, err := c.post(url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var lockTransaction models.LockTransaction
	err = json.NewDecoder(resp.Body).Decode(&lockTransaction)
	if err != nil {
		return nil, err
	}
	return &lockTransaction, nil
}

func (c *Client) StateOfVestingAssets(beneficiary string) (*[]models.VestingState, error) {
	url := fmt.Sprintf("/contracts/vesting/state/%s", beneficiary)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var vestingState []models.VestingState
	err = json.NewDecoder(resp.Body).Decode(&vestingState)
	if err != nil {
		return nil, err
	}
	return &vestingState, nil

}

func (c *Client) CollectAssets(beneficiary string) (*models.CollectTransaction, error) {
	url := fmt.Sprintf("/contracts/vesting/collect/%s", beneficiary)
	resp, err := c.post(url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var collectTransaction models.CollectTransaction
	err = json.NewDecoder(resp.Body).Decode(&collectTransaction)
	if err != nil {
		return nil, err
	}
	return &collectTransaction, nil
}
