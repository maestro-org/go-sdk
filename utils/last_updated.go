package utils

type LastUpdated struct {
	BlockHash string `json:"block_hash"`
	BlockSlot int64  `json:"block_slot"`
	Timestamp string `json:"timestamp"`
}
