package main

import lru "github.com/hashicorp/golang-lru"

type ARC struct {
	v *lru.ARCCache
}

func NewARC(size int) Cache {
	cache, err := lru.NewARC(size)
	if err != nil {
		panic(err)
	}
	return &ARC{
		v: cache,
	}
}

func (c *ARC) Name() string {
	return "arc"
}

func (c *ARC) Set(key string) {
	c.v.Add(key, key)
}

func (c *ARC) Get(key string) bool {
	_, ok := c.v.Get(key)
	return ok
}
