package common

import (
	"sync"
	"time"
)

var Cache *CommonCache

func init() {
	Cache = &CommonCache{}
	Cache.data = make(map[string]interface{})
}

type CommonCache struct {
	lock sync.RWMutex
	data map[string]interface{}
}

// get cache value by key
func (c *CommonCache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	d, ok := c.data[key]
	return d, ok
}

// set cache value, if minutes is 0 - caching forever
func (c *CommonCache) Set(key string, value interface{}, minutes int) {
	if value == nil {
		return
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = value
	if minutes != 0 {
		go time.AfterFunc(time.Duration(minutes)*time.Minute, func() { c.Forget(key) })
	}
}

// clear cache value by key
func (c *CommonCache) Forget(key string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, key)

	return !c.Has(key)
}

// check cache key is exists
func (c *CommonCache) Has(key string) bool {
	return c.data[key] != nil
}
