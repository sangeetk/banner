package banner

import (
	"fmt"
	"time"
)

// Banner image
type Banner struct {
	ID       string
	ActiveAt int64
	ExpireAt int64
	Image    string
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
