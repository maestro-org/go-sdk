package models

type LockTransaction struct {
	CborHex string `json:"cbor_hex"`
	TxHash  string `json:"tx_hash"`
}

type VestingState struct {
	AssetName                string `json:"asset_name"`
	AssetSymbol              string `json:"asset_symbol"`
	RemainingVestingQuantity int64  `json:"remaining_vesting_quantity"`
	TotalInstallments        int64  `json:"total_installments"`
	TotalVestingQuantity     int64  `json:"total_vesting_quantity"`
	VestingPeriodStart       int64  `json:"vesting_period_start"`
	VestingPeriodEnd         int64  `json:"vesting_period_end"`
}

type CollectTransaction struct {
	CborHex string `json:"cbor_hex"`
	TxHash  string `json:"tx_hash"`
}
