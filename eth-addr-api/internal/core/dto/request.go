package dto

import "app/internal/core/domains"

type AddAddressToWatchRequest struct {
	domains.Address
	FromBlock int `json:"fromBlock"`
	ToBlock int `json:"toBlock"`
}