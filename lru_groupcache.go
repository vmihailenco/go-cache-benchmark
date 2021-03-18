package main

import (
	"sync"

	"github.com/golang/groupcache/lru"
)

type GroupCacheLRU struct {
	mu sync.Mutex
	v  *lru.Cache
}

func NewGroupCacheLRU(size int) Cache {
	return &GroupCacheLRU{
		v: lru.New(size),
	}
}

func (c *GroupCacheLRU) Name() string {
	return "lru-groupcache"
}

func (c *GroupCacheLRU) Set(key string) {
	c.mu.Lock()
	c.v.Add(key, key)
	c.mu.Unlock()
}

func (c *GroupCacheLRU) Get(key string) bool {
	c.mu.Lock()
	_, ok := c.v.Get(key)
	c.mu.Unlock()
	return ok
}

func (c *GroupCacheLRU) Close() {}
