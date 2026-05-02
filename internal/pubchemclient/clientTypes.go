package pubchemclient

import (
	"net/http"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/cache"
)

type PubChemClient struct {
	cache      *cache.Cache
	httpClient http.Client
}
