package cache

type Cache struct {
}

func New() {

}

// Put adds an item to the cache. If already present, the item with
// the specified key is overwritten.
// If the optional duration (in seconds) parameter is supplied, the
// item will last for that duration, else the item would remain in the cache
// until it's manually removed.
func (c *Cache) Put(key string, value any, duration ...int) {

}

// Add adds an item to the cache if the item does not already exists.
// If the optional duration (in seconds) parameter is supplied, the
// item will last for that duration, else the item would remain in the cache
// until it's manually removed.
// Returns a boolean if the item was added.
func (c *Cache) Add(key string, value any, duration ...int) bool {
	return false
}

// Forever is the same as Put, except that the item remains in the cache
// until manually removed.
func (c *Cache) Forever(key string, value any) {

}

// Get returns the element with the specified key from the cache.
// optional defaultfunc parameter
func (c *Cache) Get(key string, defaultFunc ...func()) any {
	return nil
}

// Has returns true if an item with the specified key is present in the cache.
// Otherwise, false is returned
func (c *Cache) Has(key string) bool {
	return false
}

// duration in seconds
func (c *Cache) Remember(key string, duration int, defaultFunc func()) any {
	return nil
}

func (c *Cache) RememberForever(key string, defaultFunc func()) any {
	return nil
}

// Pull retrieves an item from the cache, deletes the item,
// and returns the item to the caller.
func (c *Cache) Pull(key string) any {
	return nil
}

// Remove deletes an item with the specified key from the cache.
func (c *Cache) Remove(key string) {

}

// Forget is an alias of remove
func (c *Cache) Forget(key string) {
	c.Remove(key)
}

// Flush removes everything from the cache
func (c *Cache) Flush() {

}
