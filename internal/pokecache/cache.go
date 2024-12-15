package pokecache

import (
	"fmt"
	"sync"
	"time"
)

// Cache -
type Cache struct {
	locationAreaCache map[string]cacheEntry
	mux               *sync.Mutex
	locationCache     map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	c := Cache{
		locationAreaCache: make(map[string]cacheEntry),
		mux:               &sync.Mutex{},
		locationCache:     make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

// Add -
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.locationAreaCache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}
func (c *Cache) AddLocation(key string, value []byte) {
	fmt.Printf("adding the %s to the cache",key )
	c.mux.Lock()
	defer c.mux.Unlock()
	c.locationCache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}

}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.locationAreaCache[key]
	
	
	return val.val, ok
}
func (c *Cache) GetLocation(location string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	
	val, ok := c.locationCache[location]
	if ok {
		fmt.Printf("%s was in the cache" , location)
	}
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.locationAreaCache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.locationAreaCache, k)
		}
	}
}
