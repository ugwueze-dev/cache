package cache

import "github.com/ugwueze-dev/cache/store"

type Cache struct {
	store store.Store
}

// TODO
/**type MultiCache struct {
	stores map[string]Cache
}**/

func New(store store.Store) *Cache {
	return &Cache{
		store: store,
	}
}

// Put adds an item to the cache. If already present, the item with
// the specified key is overwritten.
// If the optional duration (in seconds) parameter is supplied, the
// item will last for that duration, else the item would remain in the cache
// until it's manually removed.
func (c *Cache) Put(key string, value any, duration ...int) {
	dur := 0
	if len(duration) > 0 {
		dur = duration[0]
	}

	c.store.Put(key, value, dur)
}

// Add adds an item to the cache if the item does not already exists.
// If the optional duration (in seconds) parameter is supplied, the
// item will last for that duration, else the item would remain in the cache
// until it's manually removed.
// Returns a boolean if the item was added.
func (c *Cache) Add(key string, value any, duration ...int) bool {
	if c.Has(key) {
		return false
	}

	c.Put(key, value, duration...)
	return true
}

// Forever is the same as Put, except that the item remains in the cache
// until manually removed.
func (c *Cache) Forever(key string, value any) {
	c.Put(key, value, 0)
}

// Get returns the element with the specified key from the cache.
// optional defaultfunc parameter
func (c *Cache) Get(key string, defaultFunc ...func() any) any {
	value := c.store.Get(key)
	if value != nil {
		return value
	}

	if len(defaultFunc) == 0 {
		return nil
	}

	return defaultFunc[0]()
}

// Has returns true if an item with the specified key is present in the cache.
// Otherwise, false is returned
func (c *Cache) Has(key string) bool {
	return c.store.Has(key)
}

// duration in seconds
func (c *Cache) Remember(key string, duration int, defaultFunc func() any) any {
	value := c.store.Get(key)
	if value != nil {
		return value
	}

	defaultValue := defaultFunc()
	c.Put(key, defaultValue, []int{duration}...)

	return defaultValue
}

func (c *Cache) RememberForever(key string, defaultFunc func() any) any {
	return c.Remember(key, 0, defaultFunc)
}

// Pull retrieves an item from the cache, deletes the item,
// and returns the item to the caller.
func (c *Cache) Pull(key string) any {
	value := c.Get(key)
	c.Remove(key)
	return value
}

// Remove deletes an item with the specified key from the cache.
func (c *Cache) Remove(key string) {
	c.store.Remove(key)
}

// Forget is an alias of remove
func (c *Cache) Forget(key string) {
	c.Remove(key)
}

// Flush removes everything from the cache
func (c *Cache) Flush() {
	c.store.Flush()
}

// Close closes the cache. Call this during application shutdown
func (c *Cache) Close() {
	c.store.Close()
}
