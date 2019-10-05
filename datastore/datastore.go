package datastore

const (
	ErrorNotFound   = "Not found"
	ErrorInvalidKey = "Invalid key"
)

type DataStore interface {
	Put(key string, data interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
}
