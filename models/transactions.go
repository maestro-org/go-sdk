package models

import "github.com/maestro-org/go-sdk/utils"

type Certificates struct {
	MirTransfers               []any `json:"mir_transfers"`
	PoolRegistrations          []any `json:"pool_registrations"`
	PoolRetirements            []any `json:"pool_retirements"`
	StakeDelegations           []any `json:"stake_delegations"`
	StakeDeregistrations       []any `json:"stake_deregistrations"`
	StakeRegistrationsReserves []any `json:"stake_registrations_reserves"`
}

type Redeemers struct {
	Certificates []any `json:"certificates"`
	Mints        []any `json:"mints"`
	Spends       []any `json:"spends"`
	Withdrawals  []any `json:"withdrawals"`
}

type TransactionDetail struct {
	AdditionalSigners []string     `json:"additional_signers"`
	BlockAbsoluteSlot int64        `json:"block_absolute_slot"`
	BlockHash         string       `json:"block_hash"`
	BlockHeight       int64        `json:"block_height"`
	BlockTimestamp    int64        `json:"block_timestamp"`
	BlockTxIndex      int64        `json:"block_tx_index"`
	Certificates      Certificates `json:"certificates"`
	CollateralInputs  []Utxo       `json:"collateral_inputs"`
	CollateralReturn  any          `json:"collateral_return"`
	Deposit           int64        `json:"deposit"`
	Fee               int64        `json:"fee"`
	Inputs            []Utxo       `json:"inputs"`
	InvalidBefore     int64        `json:"invalid_before"`
	InvalidHereafter  int64        `json:"invalid_hereafter"`
	Metadata          any          `json:"metadata"`
	Mint              []any        `json:"mint"`
	Outputs           []Utxo       `json:"outputs"`
	Redeemers         []Redeemers  `json:"redeemers"`
	ReferenceInputs   []any        `json:"reference_inputs"`
	ScriptsExecuted   []Script     `json:"scripts_executed"`
	ScriptsSuccesful  bool         `json:"scripts_succesful"`
	Size              int64        `json:"size"`
	TxHash            string       `json:"tx_hash"`
	Withdrawals       []any        `json:"withdrawals"`
}

type TransactionDetails struct {
	Data        TransactionDetail `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type TransactionOutputFromReference struct {
	Data        Utxo              `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type TxoReference struct {
	TxHash string `json:"tx_hash"`
	Index  int    `json:"index"`
}

type TransactionOutputsFromReferences struct {
	Data        []Utxo            `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}
