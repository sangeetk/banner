package banner

import (
	"fmt"
	"testing"

	"github.com/sangeetk/banner"
)

func TestScheduler(t *testing.T) {
	fmt.Println("Testing scheduling of banners...")

	// Create a new scheduler
	var s = New()
	s.Debug()

	if err := s.Schedule(banner.Banner{ID: "A", ActiveAt: 20, ExpireAt: 40}); err != nil {
		fmt.Println(err)
	}
	s.Debug()

	if err := s.Schedule(banner.Banner{ID: "B", ActiveAt: 10, ExpireAt: 30}); err != nil {
		fmt.Println(err)
	}
	s.Debug()

}
