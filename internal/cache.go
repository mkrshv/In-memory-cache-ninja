package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	storage map[string]*Item
	sync    sync.RWMutex
}

type Item struct {
	value any
	ttl   time.Duration
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]*Item),
	}
}

func (c *Cache) set(key string, value any, ttl time.Duration) {
	c.sync.Lock()
	defer c.sync.Unlock()

	c.storage[key] = &Item{
		value: value,
		ttl:   ttl,
	}

	go func() {
		<-time.After(ttl)
		c.Delete(key)
	}()

}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.set(key, value, ttl)
}

func (c *Cache) Get(key string) (any, error) {
	if v, ok := c.storage[key]; !ok {
		return nil, errors.New("item not found")
	} else {
		return v.value, nil
	}
}

func (c *Cache) Delete(key string) {
	c.sync.RLock()
	defer c.sync.RUnlock()

	delete(c.storage, key)
}

// Set(key string, value interface{}), Get(key string), Delete(key), cache.New()
