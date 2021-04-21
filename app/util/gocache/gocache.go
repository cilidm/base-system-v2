package gocache

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

var (
	cae  *cache.Cache
	once sync.Once
)

func Instance() *cache.Cache {
	once.Do(func() {
		cae = cache.New(3*time.Minute, 10*time.Minute)
	})
	return cae
}
