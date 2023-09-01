package client

import (
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/utils"
)

type OperationalCertificate struct {
	HotVkey        string `json:"hot_vkey"`
	KesPeriod      int64  `json:"kes_period"`
	KesSignature   string `json:"kes_signature"`
	SequenceNumber int64  `json:"sequence_number"`
}

type TotalExUnits struct {
	Mem   int64 `json:"mem"`
	Steps int64 `json:"steps"`
}

type BlockInfo struct {
	Data struct {
		AbsoluteSlot           int64                  `json:"absolute_slot"`
		BlockProducer          string                 `json:"block_producer"`
		Confirmations          int64                  `json:"confirmations"`
		Epoch                  int64                  `json:"epoch"`
		EpochSlot              int64                  `json:"epoch_slot"`
		Era                    string                 `json:"era"`
		Hash                   string                 `json:"hash"`
		Height                 string                 `json:"height"`
		OperationalCertificate OperationalCertificate `json:"operational_certificate"`
		PreviousBlock          string                 `json:"previous_block"`
		ProtocolVersion        []int                  `json:"protocol_version"`
		ScriptInvocations      int32                  `json:"script_invocations"`
		Size                   int32                  `json:"size"`
		Timestamp              string                 `json:"timestamp"`
		TotalExUnits           TotalExUnits           `json:"total_ex_units"`
		TotalFees              int64                  `json:"total_fees"`
		TotalOutputLovelace    []string               `json:"total_output_lovelace"`
		TxHashes               []string               `json:"tx_hashes"`
		VrfKey                 string                 `json:"vrf_key"`
	} `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}

func (c *Client) BlockInfo(blockHeight int64) (*BlockInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blocks/%d", c.baseUrl, blockHeight), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res := BlockInfo{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
