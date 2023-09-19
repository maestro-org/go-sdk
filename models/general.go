package models

import "github.com/maestro-org/go-sdk/utils"

type ChainTipData struct {
	BlockHash string `json:"block_hash"`
	Height    int64  `json:"height"`
	Slot      int64  `json:"slot"`
}

type ChainTip struct {
	Data        ChainTipData      `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type EraTimeStamp struct {
	Epoch int64 `json:"epoch"`
	Slot  int64 `json:"slot"`
	Time  int64 `json:"time"`
}

type EraParams struct {
	EpochLength int64 `json:"epoch_length"`
	SafeZone    int64 `json:"safe_zone"`
	SlotLength  int64 `json:"slot_length"`
}

type Era struct {
	End        EraTimeStamp `json:"end"`
	Parameters EraParams    `json:"parameters"`
	Start      EraTimeStamp `json:"start"`
}

type EraHistory struct {
	Data        []Era             `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type ProtocolParameters struct {
	Data        ProtocolParams    `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type ExUnits struct {
	Memory int64 `json:"memory"`
	Steps  int64 `json:"steps"`
}

type StringExUnits struct {
	Memory string `json:"memory"`
	Steps  string `json:"steps"`
}

type ProtocolVersion struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
}

type ProtocolParams struct {
	CoinsPerUtxoByte                int64           `json:"coins_per_utxo_byte"`
	CollateralPercentage            int64           `json:"collateral_percentage"`
	CostModels                      any             `json:"cost_models"`
	DesiredNumberOfPools            int64           `json:"desired_number_of_pools"`
	MaxBlockBodySize                int64           `json:"max_block_body_size"`
	MaxBlockHeaderSize              int64           `json:"max_block_header_size"`
	MaxCollateralInputs             int64           `json:"max_collateral_inputs"`
	MaxExecutionUnitsPerBlock       ExUnits         `json:"max_execution_units_per_block"`
	MaxExecutionUnitsPerTransaction ExUnits         `json:"max_execution_units_per_transaction"`
	MaxTxSize                       int64           `json:"max_tx_size"`
	MaxValueSize                    int64           `json:"max_value_size"`
	MinFeeCoefficient               int64           `json:"min_fee_coefficient"`
	MinFeeConstant                  int64           `json:"min_fee_constant"`
	MinPoolCost                     int64           `json:"min_pool_cost"`
	MonetaryExpansion               string          `json:"monetary_expansion"`
	PoolDeposit                     int64           `json:"pool_deposit"`
	PoolInfluence                   string          `json:"pool_influence"`
	PoolRetirementEpochBound        int64           `json:"pool_retirement_epoch_bound"`
	Prices                          StringExUnits   `json:"prices"`
	ProtocolVersion                 ProtocolVersion `json:"protocol_version"`
	StakeKeyDeposit                 int64           `json:"stake_key_deposit"`
	TreasuryExpansion               string          `json:"treasury_expansion"`
}
