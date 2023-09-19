package models

import "github.com/maestro-org/go-sdk/utils"

type Epoch struct {
	BlkCount  int `json:"blk_count"`
	EpochNo   int `json:"epoch_no"`
	Fees      int `json:"fees"`
	StartTime int `json:"start_time"`
	TxCount   int `json:"tx_count"`
}

type EpochResp struct {
	Data        Epoch             `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}
