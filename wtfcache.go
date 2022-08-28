package main

import (
	"github.com/irr123/wtfcache"
)

type WTFCache struct {
	v *wtfcache.LRUWithLock[string, string]
}

func NewWTFCache(size int) Cache {
	return &WTFCache{
		v: wtfcache.New[string, string]().MakeWithLock(size),
	}
}

func (c *WTFCache) Name() string {
	return "wtfcache"
}

func (c *WTFCache) Set(key string) {
	c.v.Set(key, key)
}

func (c *WTFCache) Get(key string) bool {
	_, ok := c.v.Get(key)
	return ok
}

func (c *WTFCache) Close() {}
