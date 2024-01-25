package store

type Store interface {
	Put(key string, data any, duration int)
	Has(key string) bool
	Get(key string) any
	Remove(key string)
	Flush()
}
