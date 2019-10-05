package banner

type DataStore interface {
	Put(key string, b Banner) error
	Get(key string) (Banner, error)
	Del(key string) error
}
