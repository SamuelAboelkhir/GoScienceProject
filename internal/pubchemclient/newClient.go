// Package pubchemclient: This package handles all the API requests to PubChem
package pubchemclient

import (
	"net/http"
	"time"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/cache"
)

func NewClient(timeout, cacheTimeout time.Duration) *PubChemClient {
	return &PubChemClient{
		cache: cache.NewCache(cacheTimeout),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
