package go_cache

import (
	"fmt"
	"time"
)

type Cache struct {
	store    map[string]Entry
	lifetime time.Duration
}

type Entry struct {
	data      []byte
	timestamp time.Time
}

func ToEntry(data []byte) Entry {
	return Entry{data, time.Now()}
}

// ClearExpired removes expired entries from the cache, it runs continuously
// at a set interval, and should be used as a background go func.
func (c *Cache) ClearExpired(interval time.Duration) {
	for range time.Tick(interval) {
		for key, entry := range c.store {
			if entry.HasExpired(c.lifetime) {
				delete(c.store, key)
			}
		}
	}
}

// Check if an Entry has expired based on the Cache lifetime
func (e Entry) HasExpired(lifetime time.Duration) bool {
	return time.Now().After(e.timestamp.Add(lifetime))
}

// Check if an Entry is fresh based on the Cache lifetime
func (e Entry) IsFresh(lifetime time.Duration) bool {
	return !e.HasExpired(lifetime)
}

// Retrieve an Entry from cache if it's present and fresh
func (c *Cache) Get(key string) []byte {
	entry, exists := c.store[key]
	if !exists {
		return nil
	}
	if entry.HasExpired(c.lifetime) {
		c.Delete(key)
		return nil
	}
	return entry.data
}

// Add object to cache
func (c *Cache) Add(key string, data []byte) {
	entry := Entry{data, time.Now()}
	c.store[key] = entry
}

// Delete object from cache
func (c *Cache) Delete(key string) {
	delete(c.store, key)
}

// Check if object is in cache
func (c *Cache) Exists(key string) bool {
	_, exists := c.store[key]
	return exists
}

// Check if object in cache is fresh
func (c *Cache) IsFresh(key string) (bool, error) {
	entry, exists := c.store[key]
	if !exists {
		return false, fmt.Errorf("%s does not exist", key)
	}
	return entry.IsFresh(c.lifetime), nil
}

func MakeCache(lifetimeSeconds int64) Cache {
	lt := time.Duration(time.Second.Nanoseconds() * lifetimeSeconds)
	return Cache{make(map[string]Entry), lt}
}
