package main

import (
	"github.com/qianbin/directcache"
)

type DirectCache struct {
	v *directcache.Cache
}

func NewDirectCache(size int) Cache {
	entrySize := 4 + 8 + 8 // hdr + klen + vlen
	c := directcache.New(entrySize * size)
	return &DirectCache{
		v: c,
	}
}

func (c *DirectCache) Name() string {
	return "directcache"
}

func (c *DirectCache) Set(key string) {
	c.v.Set([]byte(key), []byte(key))
}

func (c *DirectCache) Get(key string) bool {
	_, ok := c.v.Get([]byte(key))
	return ok
}

func (c *DirectCache) Close() {}
