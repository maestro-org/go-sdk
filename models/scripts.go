package models

import "github.com/maestro-org/go-sdk/utils"

type ScriptVersion string

const (
	PlutusV1 ScriptVersion = "plutusv1"
	PlutusV2 ScriptVersion = "plutusv2"
)

type Script struct {
	Bytes string        `json:"bytes"`
	Hash  string        `json:"hash"`
	Json  any           `json:"json"`
	Type  ScriptVersion `json:"type"`
}
type ScriptByHash struct {
	Data        Script            `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}
