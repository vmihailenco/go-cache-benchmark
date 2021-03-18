package main

import (
	"go-cache-benchmark/slru"
	"sync"
)

type SLRU struct {
	mu sync.Mutex
	v  *slru.Cache
}

func NewSLRU(size int) Cache {
	return &SLRU{
		v: slru.New(int(float64(size)*0.2), int(float64(size)*0.8)),
	}
}

func (c *SLRU) Name() string {
	return "slru"
}

func (c *SLRU) Set(key string) {
	c.mu.Lock()
	c.v.Set(key, key)
	c.mu.Unlock()
}

func (c *SLRU) Get(key string) bool {
	c.mu.Lock()
	v := c.v.Get(key)
	c.mu.Unlock()
	return v != nil
}

func (c *SLRU) Close() {}
