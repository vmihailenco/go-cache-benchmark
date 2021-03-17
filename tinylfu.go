package main

import (
	"sync"

	"github.com/dgryski/go-tinylfu"
)

type TinyLFU struct {
	mu sync.Mutex
	v  *tinylfu.T
}

func NewTinyLFU(size int) Cache {
	return &TinyLFU{
		v: tinylfu.New(size, size*10),
	}
}

func (c *TinyLFU) Name() string {
	return "tinylfu"
}

func (c *TinyLFU) Set(key string) {
	c.mu.Lock()
	c.v.Add(key, key)
	c.mu.Unlock()
}

func (c *TinyLFU) Get(key string) bool {
	c.mu.Lock()
	_, ok := c.v.Get(key)
	c.mu.Unlock()
	return ok
}
