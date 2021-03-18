package main

import lru "github.com/hashicorp/golang-lru"

type HashicorpLRU struct {
	v *lru.Cache
}

func NewHashicorpLRU(size int) Cache {
	cache, err := lru.New(size)
	if err != nil {
		panic(err)
	}
	return &HashicorpLRU{
		v: cache,
	}
}

func (c *HashicorpLRU) Name() string {
	return "lru-hashicorp"
}

func (c *HashicorpLRU) Set(key string) {
	c.v.Add(key, key)
}

func (c *HashicorpLRU) Get(key string) bool {
	_, ok := c.v.Get(key)
	return ok
}

func (c *HashicorpLRU) Close() {}
