package main

import (
	"sync"

	"github.com/dgryski/go-s4lru"
)

type S4LRU struct {
	mu sync.Mutex
	v  *s4lru.Cache
}

func NewS4LRU(size int) Cache {
	return &S4LRU{
		v: s4lru.New(size),
	}
}

func (c *S4LRU) Name() string {
	return "s4lru"
}

func (c *S4LRU) Set(key string) {
	c.mu.Lock()
	c.v.Set(key, key)
	c.mu.Unlock()
}

func (c *S4LRU) Get(key string) bool {
	c.mu.Lock()
	_, ok := c.v.Get(key)
	c.mu.Unlock()
	return ok
}

func (c *S4LRU) Close() {}
