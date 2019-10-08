package banner

import (
	"image"
	"path/filepath"
	"time"
)

// Common errors
const (
	ErrorNotFound        = "Not found"
	ErrorInvalidSchedule = "Invalid schedule"
	ErrorExpiredSchedule = "Expired schedule"
)

// Banner image
type Banner struct {
	ID          string `json:"id"`
	Filename    string `json:"filename"`
	ActiveAt    int64  `json:"active_at"`
	ExpireAt    int64  `json:"expire_at"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}

// NewFile creates a banner using the file on disk
func NewFile(id, path string) (*Banner, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	b := &Banner{
		ID:       id,
		Filename: filepath.Base(path),
		Image:    content,
	}
	return b, nil
}

// NewURL create a banner from the given URL
func NewURL(id, url string) (*Banner, error) {
	// create a URL from string
	fileURL, err := url.Parse(url)
	if err != nil {
		return nil, err
	}
	// Get filename
	segments := strings.Split(fileURL.Path, "/")
	fileName = segments[len(segments)-1]

	// Get the file fom the given url
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	b := &Banner{
		ID:       id,
		Filename: fileName,
		Image:    content,
	}

	return b, nil
}

// AddDescription adds detail about banner
func (b *Banner) AddDescription(s string) *Banner {
	b.Description = s
	return b
}

// AddDuration sets active duration of the banner
func (b *Banner) AddDuration(start, end time.Time) *Banner {
	b.ActiveAt = start.Unix()
	b.ExpireAt = end.Unix()
	return b
}

// Delete the banner and free storage
func (b *Banner) Delete() error {
	delete(b.Image)
	return nil
}

// IsActive tells if the banner is active or expired
func (b *Banner) IsActive() bool {
	return b.ActiveAt <= time.Now() && b.ExpireAt > time.Now()
}
