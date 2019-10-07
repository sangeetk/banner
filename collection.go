package banner

import (
	"time"
)

// Collection of active banners
type Collection struct {
	ID        string
	DataStore *DataStore
}

// NewCollection using the given datatore
func NewCollection(id string, ds *DataStore) *Collection {
	return &Collection{ID: id, DataStore: ds}
}

// GetActive returns an active banner with nearest expiry
func (c *Collection) GetActive() (*Banner, error) {
	return nil, nil
}

// GetAllActive returns all active banners currently active
func (c *Collection) GetAllActive() ([]*Banner, error) {
	return nil, nil
}

// Add a banner to the collection
func (c *Collection) Add(b *Banner) error {
	return c.DataStore.Put(b.ID, b)
}

// Remove a banner from the collection
func (c *Collection) Remove(b *Banner) error {
	return nil
}

// PreviewByID returns banner by id, ignoring its scheduled duration
func (c *Collection) PreviewByID(id string) *Banner {
	return nil
}

// PreviewByTime returns banner that is scheduled in future time
func (c *Collection) PreviewByTime(t time.Time) *Banner {
	return nil
}

// PreviewAll returns all banners that are scheduled in future
func (c *Collection) PreviewAll() []*Banner {
	return []*Banner{}
}
