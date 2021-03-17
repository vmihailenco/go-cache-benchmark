package main

import "github.com/dgryski/go-tinylfu"

type TinyLFU struct {
	v *tinylfu.T
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
	c.v.Add(key, key)
}

func (c *TinyLFU) Get(key string) bool {
	_, ok := c.v.Get(key)
	return ok
}
