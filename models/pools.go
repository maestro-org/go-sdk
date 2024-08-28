package models

import "github.com/maestro-org/go-sdk/utils"

type Pool struct {
	PoolIdBech32 string `json:"pool_id_bech32"`
	Ticker       string `json:"ticker"`
}

type RegisteredPools struct {
	Data        []Pool            `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type Block struct {
	AbsSlot     int64  `json:"abs_slot"`
	BlockHash   string `json:"block_hash"`
	BlockHeight int64  `json:"block_height"`
	BlockTime   int64  `json:"block_time"`
	EpochNo     int64  `json:"epoch_no"`
	EpochSlot   int64  `json:"epoch_slot"`
}

type PoolMintedBlocks struct {
	Data        []Block           `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
	NextCursor  string            `json:"next_cursor"`
}

type StakePoolDelegator struct {
	Amount                 string `json:"amount"`
	LatestDelegationTxHash string `json:"latest_delegation_tx_hash"`
	StakeAddress           string `json:"stake_address"`
}

type StakePoolDelegators struct {
	Data        []StakePoolDelegator `json:"data"`
	LastUpdated utils.LastUpdated    `json:"last_updated"`
	NextCursor  string               `json:"next_cursor"`
}

type StakePoolHistoryData struct {
	ActiveStake    int64  `json:"active_stake"`
	ActiveStakePct string `json:"active_stake_pct"`
	BlockCnt       int64  `json:"block_cnt"`
	DelegRewards   int64  `json:"deleg_rewards"`
	DelegatorCnt   int64  `json:"delegator_cnt"`
	EpochNo        int64  `json:"epoch_no"`
	EpochRos       string `json:"epoch_ros"`
	FixedCost      int64  `json:"fixed_cost"`
	Margin         int64  `json:"margin"`
	PoolFees       int64  `json:"pool_fees"`
	SaturationPct  any    `json:"saturation_pct"`
}
type StakePoolHistory struct {
	Data        []StakePoolHistoryData `json:"data"`
	LastUpdated utils.LastUpdated      `json:"last_updated"`
	NextCursor  string                 `json:"next_cursor"`
}

type Relay struct {
	Dns  string `json:"dns"`
	Ipv4 string `json:"ipv4"`
	Ipv6 string `json:"ipv6"`
	Port int64  `json:"port"`
	Srv  string `json:"srv"`
}

type StakePoolDetails struct {
	ActiveStake    int64    `json:"active_stake"`
	BlockCount     int64    `json:"block_count"`
	FixedCost      int64    `json:"fixed_cost"`
	LiveDelegators int64    `json:"live_delegators"`
	LivePledge     int64    `json:"live_pledge"`
	LiveSaturation string   `json:"live_saturation"`
	LiveStake      int64    `json:"live_stake"`
	Margin         int64    `json:"margin"`
	MetaHash       any      `json:"meta_hash"`
	MetaJson       any      `json:"meta_json"`
	MetaUrl        any      `json:"meta_url"`
	OpCert         string   `json:"op_cert"`
	OpCertCounter  int64    `json:"op_cert_counter"`
	Owners         []string `json:"owners"`
	Pledge         int64    `json:"pledge"`
	PoolIdBech32   string   `json:"pool_id_bech32"`
	PoolIdHex      string   `json:"pool_id_hex"`
	PoolStatus     string   `json:"pool_status"`
	Relays         []Relay  `json:"relays"`
	RetiringEpoch  any      `json:"retiring_epoch"`
	RewardAddr     string   `json:"reward_addr"`
	Sigma          string   `json:"sigma"`
	VrfKeyHash     string   `json:"vrf_key_hash"`
}

type StakePoolInformation struct {
	Data        StakePoolDetails  `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type Poolmetadata struct {
	MetaHash     string `json:"meta_hash"`
	MetaJson     any    `json:"meta_json"`
	MetaUrl      string `json:"meta_url"`
	PoolIdBech32 string `json:"pool_id_bech32"`
}

type StakePoolMetadata struct {
	Data        Poolmetadata      `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type RelaysAndId struct {
	PoolIdBech32 string  `json:"pool_id_bech32"`
	Relays       []Relay `json:"relays"`
}

type StakePoolRelays struct {
	Data        []RelaysAndId     `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

type StakePoolUpdates struct {
	Data        []StakePoolDetails `json:"data"`
	LastUpdated utils.LastUpdated  `json:"last_updated"`
}
