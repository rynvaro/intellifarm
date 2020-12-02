package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var L1Cache = cache.New(5*time.Minute, 10*time.Minute)
