package providers

type Store interface {
	Put(key string, data any)
	Get(key string) any
	Remove(key string)
	Flush()
}
