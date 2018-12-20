package gohelper

import (
	"sync"
	"time"
)

type cacheBucket struct {
	ele     interface{}
	timeout time.Time
}

// GcInterval: the interval of gc in Seconds
// GcCallback: the function will be called after gc
type LocalCacheOption struct {
	GcInterval time.Duration
	GcCallback func(lc *LocalCache)
}

type LocalCache struct {
	buckets map[string]cacheBucket
	mu      *sync.RWMutex
	opt     *LocalCacheOption
}

func NewLocalCache(opt *LocalCacheOption) (lc *LocalCache) {
	buckets := make(map[string]cacheBucket)
	mu := new(sync.RWMutex)

	lc = &LocalCache{buckets, mu, opt}

	go intervalGc(lc)

	return
}

// ttl: the expiration time in Seconds
func (lc *LocalCache) Set(key string, val interface{}, ttl time.Duration) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.buckets[key] = cacheBucket{
		val,
		time.Now().Add(time.Second * ttl),
	}
}

func (lc *LocalCache) Get(key string) (interface{}, bool) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	b, exist := lc.buckets[key]
	if !exist {
		return nil, false
	}

	if time.Now().After(b.timeout) {
		lc.mu.RUnlock()
		lc.Delete(key)
		lc.mu.RLock()

		return nil, false
	}

	return b.ele, true
}

func (lc *LocalCache) Delete(key string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.buckets, key)
}

func (lc *LocalCache) Length() int {
	return len(lc.buckets)
}

func (lc *LocalCache) Gc() {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	for k, v := range lc.buckets {
		if time.Now().After(v.timeout) {
			lc.mu.RUnlock()
			lc.Delete(k)
			lc.mu.RLock()
		}
	}
}

func intervalGc(lc *LocalCache) {
	ticker := time.NewTicker(lc.opt.GcInterval * time.Second)

	for {
		select {
		case <-ticker.C:
			lc.Gc()
			if nil != lc.opt.GcCallback {
				lc.opt.GcCallback(lc)
			}
		}
	}
}
