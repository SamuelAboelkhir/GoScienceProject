package client

import (
	"net/http"

	"github.com/SamuelAboelkhir/GoScienceProject/internal/cache"
)

type Client struct {
	cache      cache.Cache
	httpClient http.Client
}
