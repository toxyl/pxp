package language

import (
	"image"
	"sync"
	"time"
)

// CacheEntry represents an entry in the image cache
type CacheEntry struct {
	key   string
	value *image.NRGBA64
	added time.Time
}

type ImageCache struct {
	active bool
	items  map[string]*CacheEntry
	mutex  sync.RWMutex
}

// NewImageCache creates a new image cache with the specified maximum size
func NewImageCache() *ImageCache {
	c := &ImageCache{
		active: true,
		items:  make(map[string]*CacheEntry),
		mutex:  sync.RWMutex{},
	}
	go func() {
		for {
			c.mutex.Lock()
			for key, val := range c.items {
				if time.Since(val.added) > 2*time.Minute {
					delete(c.items, key)
				}
			}
			c.mutex.Unlock()
			time.Sleep(time.Minute)
		}
	}()
	return c
}

// Get retrieves an image from the cache
func (c *ImageCache) Get(key string) (*image.NRGBA64, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if !c.active {
		return nil, false
	}

	if elem, found := c.items[key]; found {
		return elem.value, true
	}
	return nil, false
}

// Get retrieves an image from the cache
func (c *ImageCache) UpdateTimestamp(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if !c.active {
		return false
	}

	if _, found := c.items[key]; !found {
		return false
	}
	c.items[key].added = time.Now()
	return true
}

// Put adds an image to the cache
func (c *ImageCache) Put(key string, value *image.NRGBA64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.active {
		return
	}

	// Add the new item
	c.items[key] = &CacheEntry{key: key, value: value, added: time.Now()}
}

func (c *ImageCache) Enable() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.active = true
}

func (c *ImageCache) Disable() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.active = false
}
