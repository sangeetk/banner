package banner

import (
	"log"
	"time"
)

// Collection of active banners
type Collection struct {
	ID        string
	DataStore DataStore
	Scheduler Scheduler
}

// NewCollection using the given datatore
func NewCollection(id string, ds DataStore) *Collection {
	return &Collection{
		ID:        id,
		DataStore: ds,
		Scheduler: Scheduler{},
	}
}

// RestoreCollection using the given datatore
func RestoreCollection(id string, ds DataStore) *Collection {
	col := &Collection{
		ID:        id,
		DataStore: ds,
		Scheduler: Scheduler{},
	}
	// Schedule all banners from the datastore
	for _, b := range col.DataStore.List() {
		col.Scheduler.Schedule(b)
	}
	return col
}

// GetActive returns an active banner with nearest expiry
func (c *Collection) GetActive() (Banner, error) {
	return c.Scheduler.Get()
}

// GetAllActive returns all active banners currently active
func (c *Collection) GetAllActive() ([]Banner, error) {
	return c.Scheduler.GetAll()
}

// Add a banner to the collection and scheduler
func (c *Collection) Add(b Banner) error {
	// Add the Banner to data store
	if err := c.DataStore.Put(b); err != nil {
		return err
	}

	// Schedule the banner to display at the requested time duraiton
	if err := c.Scheduler.Schedule(b); err != nil {
		// Remove the banner from data store
		c.DataStore.Del(b.ID)
		return err
	}
	return nil
}

// Remove a banner from the collection
func (c *Collection) Remove(b Banner) error {
	if err := c.DataStore.Del(b.ID); err != nil {
		log.Fatal("Datastore error ", err)
		return err
	}
	err := c.Scheduler.Unschedule(b.ID)
	if err != nil {
		log.Fatal("Scheduler error ", err)
	}
	return err
}

// Preview returns banner that is scheduled in future time
func (c *Collection) Preview(t time.Time) (Banner, error) {
	return c.Scheduler.GetByTime(t.Unix())
}

// PreviewAll returns all banners that are scheduled in future
func (c *Collection) PreviewAll(t time.Time) ([]Banner, error) {
	return c.Scheduler.GetAllByTime(t.Unix())
}
