package models

type TxManagerState struct {
	Block           string `json:"block"`            // Block number
	State           string `json:"state"`            // Transaction state
	Timestamp       string `json:"timestamp"`        // Timestamp of last state change
	TransactionHash string `json:"transaction_hash"` // Transaction hash
}
