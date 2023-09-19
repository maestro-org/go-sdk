package models

import "github.com/maestro-org/go-sdk/utils"

type AccountAddresses struct {
	Data        []string          `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type AccountAsset struct {
	Amount int64  `json:"amount"`
	Unit   string `json:"unit"`
}

type AccountAssets struct {
	Data        []AccountAsset    `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type StakeAccountHistoryItem struct {
	ActiveStake int64  `json:"active_stake"`
	EpochNo     int64  `json:"epoch_no"`
	PoolId      string `json:"pool_id"`
}

type StakeAccountHistory struct {
	Data        []StakeAccountHistoryItem `json:"data"`
	LastUpdated utils.LastUpdated         `json:"last_updated"`
	NextCursor  string                    `json:"next_cursor"`
}

type AccountInformation struct {
	DelegatedPool    string `json:"delegated_pool"`
	Registered       bool   `json:"registered"`
	RewardsAvailable int64  `json:"rewards_available"`
	StakeAddress     string `json:"stake_address"`
	TotalBalance     int64  `json:"total_balance"`
	TotalRewarded    int64  `json:"total_rewarded"`
	TotalWithdrawn   int64  `json:"total_withdrawn"`
	UtxoBalance      int64  `json:"utxo_balance"`
}

type StakeAccountInformation struct {
	Data        AccountInformation `json:"data"`
	LastUpdated utils.LastUpdated  `json:"last_updated"`
}

type StakeRewardType string

const (
	StakeRewardMember StakeRewardType = "member"
	StakeRewardLeader StakeRewardType = "leader"
	StakeRewardRefund StakeRewardType = "refund"
)

type StakeAccountReward struct {
	Amount         int64           `json:"amount"`
	EarnedEpoch    int64           `json:"earned_epoch"`
	PoolId         string          `json:"pool_id"`
	SpendableEpoch int64           `json:"spendable_epoch"`
	Type           StakeRewardType `json:"type"`
}
type StakeAccountRewards struct {
	Data        []StakeAccountReward `json:"data"`
	LastUpdated utils.LastUpdated    `json:"last_updated"`
	NextCursor  string               `json:"next_cursor"`
}

type StakeUpdateAction string

const (
	StakeUpdateRegistration   StakeUpdateAction = "registration"
	StakeUpdateDeregistration StakeUpdateAction = "deregistration"
	StakeUpdateDelegation     StakeUpdateAction = "delegation"
	StakeUpdateWithdrawal     StakeUpdateAction = "withdrawal"
)

type StakeAccountUpdate struct {
	AbsoluteSlot int64             `json:"abs_slot"`
	Action       StakeUpdateAction `json:"action"`
	EpochNo      int64             `json:"epoch_no"`
	TxHash       string            `json:"tx_hash"`
}

type StakeAccountUpdates struct {
	Data        []StakeAccountUpdate `json:"data"`
	LastUpdated utils.LastUpdated    `json:"last_updated"`
	NextCursor  string               `json:"next_cursor"`
}
