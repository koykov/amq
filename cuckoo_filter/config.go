package cuckoo

import "github.com/koykov/amq"

const (
	defaultKicksLimit = 500
	defaultSeed       = 2077
)

type Config struct {
	Size       uint64
	KicksLimit uint64
	Hasher     amq.Hasher
	Seed       uint64
}

func (c *Config) copy() *Config {
	cpy := *c
	return &cpy
}
