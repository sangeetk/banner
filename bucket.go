package banner

import (
	"time"

	"github.com/sangeetk/banner/datastore"
)

type Bucket struct {
	id string
	ds datastore.DataStore
}

func NewBucket(id string, ds datastore.DataStore) *Bucket {
	return &Bucket{id: id, ds: ds}
}

func (b *Bucket) Put(bnr *Banner) error {
	return b.ds.Put(bnr.ID, bnr)
}

// Get an active banner, if there are more than one
// then returns the banner with nearest expiry time
func (b *Bucket) Get(tz time.Time) (*Banner, error) {
	return nil, nil
}

// Preview banner by id
func (b *Bucket) Preview(id int) *Banner {
	return nil
}

// Preview all banners scheduled in future
func (b *Bucket) PreviewAll() []*Banner {
	return []*Banner{}
}

// Create a Banner from the given file path
func (b *Bucket) NewFile(file string) (*Banner, error) {
	return nil, nil
}

// Create a Banner from the given URL
func (b *Bucket) NewURL(url string) (*Banner, error) {
	return nil, nil
}
