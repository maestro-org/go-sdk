package models

import "github.com/maestro-org/go-sdk/utils"

type BasicResponse struct {
	Data        string            `json:"data"`
	LastUpdated utils.LastUpdated `json:"last_updated"`
}
