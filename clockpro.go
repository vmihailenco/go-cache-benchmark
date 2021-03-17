package main

import (
	"sync"

	"github.com/dgryski/go-clockpro"
)

type ClockPro struct {
	mu sync.Mutex
	v  *clockpro.Cache
}

func NewClockPro(size int) Cache {
	return &ClockPro{
		v: clockpro.New(size),
	}
}

func (c *ClockPro) Name() string {
	return "clockpro"
}

func (c *ClockPro) Set(key string) {
	c.mu.Lock()
	c.v.Set(key, key)
	c.mu.Unlock()
}

func (c *ClockPro) Get(key string) bool {
	c.mu.Lock()
	v := c.v.Get(key)
	c.mu.Unlock()
	return v != nil
}
