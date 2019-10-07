package banner

const (
	ErrorNotFound = "Not found"
)

// DataStore interface
type DataStore interface {
	Put(key string, data interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
}
