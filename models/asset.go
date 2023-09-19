package models

import "github.com/maestro-org/go-sdk/utils"

type AccountHoldingAsset struct {
	Account string `json:"account"`
	Amount  int64  `json:"amount"`
}

type AccountsHoldingAsset struct {
	Data        []AccountHoldingAsset `json:"data"`
	LastUpdated utils.LastUpdated     `json:"last_updated"`
	NextCursor  string                `json:"next_cursor"`
}

type AddressHoldingAsset struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

type AddressesHoldingAsset struct {
	Data        []AddressHoldingAsset `json:"data"`
	LastUpdated utils.LastUpdated     `json:"last_updated"`
	NextCursor  string                `json:"next_cursor"`
}

type AssetInformations struct {
	Data        AssetInformation  `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type AssetTransactions struct {
	Data        []PolicyTransaction `json:"data"`
	LastUpdated utils.LastUpdated   `json:"last_updated"`
	NextCursor  string              `json:"next_cursor"`
}

type Metadata struct {
	Json any `json:"json"`
	Key  any `json:"key"`
}

type AssetUpdate struct {
	BlockTimestamp int64    `json:"block_timestamp"`
	Metadata       Metadata `json:"metadata"`
	MintAmount     int64    `json:"mint_amount"`
	TxHash         string   `json:"tx_hash"`
}

type AssetUpdates struct {
	Data        []AssetUpdate     `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type AssetUtxo struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
	Index   int64  `json:"index"`
	Slot    int64  `json:"slot"`
	TxHash  string `json:"tx_hash"`
}

type AssetUtxos struct {
	Data        []AssetUtxo       `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}
