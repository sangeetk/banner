package banner

import (
	"fmt"
	"image"
	"time"
)

// Banner image
type Banner struct {
	ID          string
	Image       image.Image
	ActiveAt    int64
	ExpireAt    int64
	Description string
}

// Get an active banner, if there are more than one
// then returns the banner with nearest expiry time
func Get(tz time.Time) (*Banner, error) {
	return nil, nil
}

// Preview banner by id
func Preview(id int) *Banner {
	return nil
}

// Preview all banners scheduled in future
func PreviewAll() []*Banner {
	return []*Banner{}
}

// Create a Banner from the given file path
func NewFile(file string) (*Banner, error) {
	return nil, nil
}

// Create a Banner from the given URL
func NewURL(url string) (*Banner, error) {
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
