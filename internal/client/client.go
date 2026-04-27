// Package client: This package handles all the API requests to PubChem
package client

import (
	"net/http"
	"time"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/cache"
)

func NewClient(timeout, cacheTimeout time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheTimeout),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
