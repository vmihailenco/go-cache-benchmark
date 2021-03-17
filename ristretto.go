package main

import "github.com/dgraph-io/ristretto"

type Ristretto struct {
	v *ristretto.Cache
}

func NewRistretto(size int) Cache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: int64(size * 10),
		MaxCost:     int64(size),
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}

	return &Ristretto{
		v: cache,
	}
}

func (c *Ristretto) Name() string {
	return "ristretto"
}

func (c *Ristretto) Set(key string) {
	c.v.Set(key, key, 1)
}

func (c *Ristretto) Get(key string) bool {
	_, ok := c.v.Get(key)
	return ok
}

func (c *Ristretto) Close() {
	c.v.Close()
}
