// Package cache: Handles caching for all the app's modules
package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	mutex    *sync.Mutex
	interval time.Duration
}
