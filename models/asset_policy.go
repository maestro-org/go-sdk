package models

import (
	"github.com/maestro-org/go-sdk/utils"
)

type AssetNameAndAmount struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}
type AccountHolding struct {
	Account string               `json:"account"`
	Assets  []AssetNameAndAmount `json:"assets"`
}

type AccountsHoldingPolicy struct {
	Data        []AccountHolding  `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type AddressHolding struct {
	Address string               `json:"address"`
	Assets  []AssetNameAndAmount `json:"assets"`
}

type AddressesHoldingPolicy struct {
	Data        []AddressHolding  `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type Standards struct {
	Cip25Metadata any `json:"cip25_metadata"`
	Cip68Metadata any `json:"cip68_metadata"`
}

type AssetInformation struct {
	AssetName             string    `json:"asset_name"`
	AssetNameAscii        string    `json:"asset_name_ascii"`
	AssetStandards        Standards `json:"asset_standards"`
	BurnTxCount           int64     `json:"burn_tx_count"`
	FingerPrint           string    `json:"fingerprint"`
	FirstMintTime         int64     `json:"first_mint_time"`
	FirstMintTx           string    `json:"first_mint_tx"`
	LatestMintTxMetadata  any       `json:"latest_mint_tx_metadata"`
	MintTxCount           int64     `json:"mint_tx_count"`
	TokenRegistryMetadata any       `json:"token_registry_metadata"`
	TotalSuppply          int64     `json:"total_supply"`
}
type PolicyInformation struct {
	Data        []AssetInformation `json:"data"`
	LastUpdated utils.LastUpdated  `json:"last_updated"`
	NextCursor  string             `json:"next_cursor"`
}

type PolicyTransaction struct {
	BlockHeight int64  `json:"block_height"`
	EpochNo     int64  `json:"epoch_no"`
	TxHash      string `json:"tx_hash"`
}

type PolicyTransactions struct {
	Data        []PolicyTransaction `json:"data"`
	LastUpdated utils.LastUpdated   `json:"last_updated"`
	NextCursor  string              `json:"next_cursor"`
}

type UtxosContainingPolicy struct {
	Data        []Utxo            `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}
