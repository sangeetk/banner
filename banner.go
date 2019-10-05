package banner

import (
	"fmt"
	"time"

	"github.com/sangeetk/banner/datastore"
)

type Bucket struct {
	id string
	ds datastore.DataStore
}

// Banner image
type Banner struct {
	ID       string
	ActiveAt int64
	ExpireAt int64
	Image    string
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

// Add scheduling to the Banner
func (b *Banner) AddDuration(start, end time.Time) *Banner {
	return nil
}

// Add description to the Banner
func (b *Banner) AddDescription(s string) *Banner {
	return nil
}

func (b *Banner) Store() (int, error) {
	return 0, nil
}

func (b *Banner) IsActive() bool {
	return true
}

func (b *Banner) String() string {
	return fmt.Sprintf("%v[%v,%v)", b.ID, b.ActiveAt, b.ExpireAt)
}
