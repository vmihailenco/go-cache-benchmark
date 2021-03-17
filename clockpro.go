package main

import "github.com/dgryski/go-clockpro"

type ClockPro struct {
	v *clockpro.Cache
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
	c.v.Set(key, key)
}

func (c *ClockPro) Get(key string) bool {
	return c.v.Get(key) != nil
}
