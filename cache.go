package cache

import (
	"time"
)

type Item struct {
	Value    string
	Deadline time.Time
}

type Cache struct {
	cacheMap map[string]Item
}

func NewCache() Cache {
	return Cache{make(map[string]Item)}
}

func (c *Cache) Get(key string) (string, bool) {
	v, ok := c.cacheMap[key]
	if !ok {
		return "", false
	}

	if v.Deadline.Before(time.Now()) {
		delete(c.cacheMap, key)
		return "", false
	}

	return v.Value, true
}

func (c *Cache) Put(key, value string) {
	c.cacheMap[key] = Item{
		Value:    value,
		Deadline: time.Now().Add(time.Minute * 15)}
}

func (c *Cache) Keys() []string {
	var keys []string

	for k := range c.cacheMap {
		keys = append(keys, k)
	}

	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.cacheMap[key] = Item{
		Value:    value,
		Deadline: deadline}
}
