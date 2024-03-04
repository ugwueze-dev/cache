MultiCache
=======

MultiCache is a Go Cache library that allows you to use cache data using multiple backends/stores simultaneously. It's inspired by the Laravel Cache library.

## Features 
Here are features of the library:

* [x] Multiple cache stores including memory, redis, or your own custom store.
* [x] Cache invalidation by expiration time.
* [x] Define default values in stores and override them when setting data.
* [] Using one or more stores at the same time.
* [] Use of Generics.

## Stores
### Memory Store
Here's how to use the memory store:

```go
import (
  "github.com/ugwueze-dev/cache"
  "github.com/ugwueze-dev/store"
)

// get memory store instance 
numItems := 1000000000  // max number of cache items
store, err := store.NewMemoryStore(numItems)
if err != nil {
  // handle error
  ...
}

cache := cache.New(store)
```

### Redis Store 
[WIP]

### Memcached Store 
[WIP]

### Filesystem Store
[WIP]

### Write your own custom store 
You can also write your own custom store by implementing the following interface:

```go
type Store interface {
	Put(key string, data any, duration int)
	Has(key string) bool
	Get(key string) any
	Remove(key string)
	Flush()
	Close()
}
```

## Usage
All of the following are available after configuring a store and obtaining a cache instance.

### Retrieving Items From the Cache
```go
...

// return a cached value when in the cache or nil when item doesn't exist
val := cache.Get("key")


// return a cached value when item exists, or a default value when item doesn't exist
val := cache.Get("key", func() any {
  return "default"
})
```

### Check if an item exists in the cache
```go
if cache.Has("key") {
  // do something
  ...
}
```

### Retrieve or store
This is used to retrieve an item from the cache when the item exists, or store a default value when the item is not present in the cache. When a default value is stored instead, it is returned to the caller.
```go
duration := 60 // 60 seconds
value := cache.Remember("key", duration, func() any {
  return "default value"
})
```
Here, If the cache contains an item with the specified key, it's retrieved. If not, the default function is invoked. The result is stored in the cache for the specified amount of time with the same key and returned.

To retrieve an item from the cache or store it forever, do the following:
```go
value := cache.RememberForever("key", func() any {
  return "default value"
})
```

### Retrieve and Delete
This retrieves an item from the cache and deletes the item from the cache. If the item is not present in the cache, `nil` is returned.
```go
value := cache.Pull(key)
```

### Storing Items in the Cache 
```go 
// store for a specified amount of time 
duration := 60 // 60 seconds 
cache.Put("key", "value", duration)

// store indefinitely
cache.Put("key", "value")
```

### Store if not present
`Add` stores an item in the cache ONLY if the item is not already present.
```go
// add for a specified amount of time 
duration := 60 // 60 seconds 
cache.Add("key", "value", duration)

// Add indefinitely
cache.Add("key", "value")
```

### Store Item Forever
`Forever` stores an item in the cache permanently.
```go
cache.Forever("key", "value")
```

### Removing Items from the Cache
`Forget` removes an item from the cache
```go
cache.Forget("key")
```

`Flush` empties the entire cache.
```go
cache.Flush()
```