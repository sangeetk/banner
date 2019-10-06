package scheduler

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

	AddBanner(s, banner.Banner{ID: "A", ActiveAt: 20, ExpireAt: 40})

	AddBanner(s, banner.Banner{ID: "B", ActiveAt: 25, ExpireAt: 30})
}

func AddBanner(s *Scheduler, b banner.Banner) {
	fmt.Println("New Banner:", b)
	if err := s.Schedule(b); err != nil {
		fmt.Println(err)
	}
	s.Debug()
}
