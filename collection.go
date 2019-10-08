package banner

import (
	"time"
)

// Collection of active banners
type Collection struct {
	ID        string
	DataStore *DataStore
	Scheduler *Scheduler
}

// NewCollection using the given datatore
func NewCollection(id string, ds *DataStore) *Collection {
	return &Collection{
		ID:        id,
		DataStore: ds,
		Scheduler: Scheduler{},
	}
}

// GetActive returns an active banner with nearest expiry
func (c *Collection) GetActive() (*Banner, error) {
	return c.Scheduler.Get()
}

// GetAllActive returns all active banners currently active
func (c *Collection) GetAllActive() ([]*Banner, error) {
	return c.Scheduler.GetAll()
}

// Add a banner to the collection and scheduler
func (c *Collection) Add(b *Banner) error {
	// Add the Banner to data store
	if err := c.DataStore.Put(b.ID, b); err != nil {
		return err
	}

	// Schedule the banner to display at the requested time duraiton
	if err := c.Scheduler.Schedule(b); err != nil {
		// Remove the banner from data store
		c.DataStore.Del(b.ID)
		return err
	}
}

// Remove a banner from the collection
func (c *Collection) Remove(b *Banner) error {
	if err := c.DataStore.Del(b.ID); err != nil {
		return err
	}
	return c.Scheduler.Unschedule(b.ID)
}

// Preview returns banner that is scheduled in future time
func (c *Collection) Preview(t time.Time) (*Banner, error) {
	return c.Scheduler.GetByTime(t.Unix())
}

// PreviewAll returns all banners that are scheduled in future
func (c *Collection) PreviewAll(t time.Time) ([]*Banner, error) {
	return c.Scheduler.GetAllByTime(t.Unix())
}
