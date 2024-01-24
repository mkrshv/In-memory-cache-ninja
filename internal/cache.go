package cache

import "sync"

type Cache struct {
	storage map[string]any
	sync    sync.RWMutex
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]any),
	}
}

func (c *Cache) set(key string, value any) {
	c.sync.Lock()
	defer c.sync.Unlock()

	c.storage[key] = value
}

func (c *Cache) Set(key string, value any) {
	c.set(key, value)
}

func (c *Cache) Get(key string) any {
	if value, ok := c.storage[key]; !ok {
		return "Element not found"
	} else {
		return value
	}
}

func (c *Cache) Delete(key string) any {
	if _, ok := c.storage[key]; !ok {
		return "Element not found"
	} else {
		delete(c.storage, key)

		return nil
	}
}

// Set(key string, value interface{}), Get(key string), Delete(key), cache.New()
