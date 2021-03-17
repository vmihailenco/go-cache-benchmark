package main

import lru "github.com/hashicorp/golang-lru"

type LRU struct {
	v *lru.Cache
}

func NewLRU(size int) Cache {
	cache, err := lru.New(size)
	if err != nil {
		panic(err)
	}
	return &LRU{
		v: cache,
	}
}

func (c *LRU) Name() string {
	return "lru"
}

func (c *LRU) Set(key string) {
	c.v.Add(key, key)
}

func (c *LRU) Get(key string) bool {
	_, ok := c.v.Get(key)
	return ok
}
