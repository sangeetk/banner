package banner

import (
	"image"
	"time"
)

// Banner image
type Banner struct {
	ID          string      `json:"id"`
	ActiveAt    int64       `json:"active_at"`
	ExpireAt    int64       `json:"expire_at"`
	Filename    string      `json:"filename"`
	Image       image.Image `json:"image"`
	Description string      `json:"description"`
}

// NewFile creates a banner using the file on disk
func NewFile(file string) (*Banner, error) {
	return nil, nil
}

// NewURL create a banner from the given URL
func NewURL(url string) (*Banner, error) {
	return nil, nil
}

// AddDescription adds detail about banner
func (b *Banner) AddDescription(s string) *Banner {
	return nil
}

// AddDuration sets active duration of the banner
func (b *Banner) AddDuration(start, end time.Time) *Banner {
	return nil
}

// Delete the banner and free storage
func (b *Banner) Delete() error {
	return nil
}

// IsActive tells if the banner is active or expired
func (b *Banner) IsActive() bool {
	return true
}
