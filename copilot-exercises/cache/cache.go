package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	cache map[string]string
	mu    sync.Mutex
}

func createCache() *Cache {
	return &Cache{
		cache: make(map[string]string),
	}
}

func (c *Cache) get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cache[key]
}

func (c *Cache) set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *Cache) delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, key)
}

func main() {
	cache := createCache()
	cache.set("key", "I am the best")
	value := cache.get("key")
	fmt.Println(value)
	cache.delete("key")
}
