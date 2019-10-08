package banner

// DataStore interface
type DataStore interface {
	Put(b Banner) error
	Get(id string) (Banner, error)
	List() []Banner
	Del(id string) error
}
