package models

import "github.com/maestro-org/go-sdk/utils"

type Datum struct {
	Bytes string `json:"bytes"`
	Json  any    `json:"json"`
}

type DatumFromHash struct {
	Data        Datum             `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}
