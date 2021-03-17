package main

import lru "github.com/hashicorp/golang-lru"

type TwoQueue struct {
	v *lru.TwoQueueCache
}

func NewTwoQueue(size int) Cache {
	cache, err := lru.New2Q(size)
	if err != nil {
		panic(err)
	}
	return &TwoQueue{
		v: cache,
	}
}

func (c *TwoQueue) Name() string {
	return "two-queue"
}

func (c *TwoQueue) Set(key string) {
	c.v.Add(key, key)
}

func (c *TwoQueue) Get(key string) bool {
	_, ok := c.v.Get(key)
	return ok
}
