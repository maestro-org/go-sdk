package models

import (
	"github.com/maestro-org/go-sdk/utils"
)

type Network string

const (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
)

type CredentialKind string

const (
	CredentialKey    CredentialKind = "key"
	CredentialScript CredentialKind = "script"
)

type PaymentCredential struct {
	Bech32 string         `json:"bech32"`
	Hex    string         `json:"hex"`
	Kind   CredentialKind `json:"kind"`
}

type StakingCredential struct {
	Bech32        string         `json:"bech32"`
	Hex           string         `json:"hex"`
	Kind          CredentialKind `json:"kind"`
	Pointer       any            `json:"pointer"`
	RewardAddress string         `json:"reward_address"`
}
type DecodedAddress struct {
	Bech32            string            `json:"bech32"`
	Hex               string            `json:"hex"`
	Network           Network           `json:"network"`
	PaymentCredential PaymentCredential `json:"payment_cred"`
	StakingCredential StakingCredential `json:"staking_cred"`
}

type AddressTransactionCount struct {
	Data        int               `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type Transaction struct {
	Input  bool   `json:"input"`
	Output bool   `json:"output"`
	Slot   int64  `json:"slot"`
	TxHash string `json:"tx_hash"`
}

type AddressTransactions struct {
	Data        []Transaction     `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type UtxoReference struct {
	Index  int64  `json:"index"`
	TxHash string `json:"tx_hash"`
}

type UtxoReferencesAtAddress struct {
	Data        []UtxoReference   `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type Asset struct {
	Amount int64  `json:"amount"`
	Unit   string `json:"unit"`
}

type Utxo struct {
	Address         string          `json:"address"`
	Assets          []Asset         `json:"assets"`
	Datum           any             `json:"datum"`
	Index           int64           `json:"index"`
	ReferenceScript ReferenceScript `json:"reference_script"`
	TxHash          string          `json:"tx_hash"`
	Slot            int64           `json:"slot"`
	TxOutCbor       string          `json:"tx_out_cbor"`
}

type ReferenceScript struct {
	Bytes string                 `json:"bytes"`
	Hash  string                 `json:"hash"`
	JSON  map[string]interface{} `json:"json"`
	Type  string                 `json:"type"`
}

type UtxosAtAddress struct {
	Data        []Utxo            `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}
