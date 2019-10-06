package scheduler

import (
	"fmt"
	"testing"
	"time"

	"github.com/sangeetk/banner"
)

func TestScheduler(t *testing.T) {
	fmt.Println("Testing scheduling of banners...")

	// Create a new scheduler
	var s = New()

	// Time here is in seconds from current time
	AddBanner(s, banner.Banner{ID: "A", ActiveAt: 10, ExpireAt: 80})
	AddBanner(s, banner.Banner{ID: "B", ActiveAt: 20, ExpireAt: 70})
	AddBanner(s, banner.Banner{ID: "C", ActiveAt: 30, ExpireAt: 60})
	AddBanner(s, banner.Banner{ID: "D", ActiveAt: 40, ExpireAt: 50})

	AddBanner(s, banner.Banner{ID: ".", ActiveAt: 0, ExpireAt: 100})

}

func AddBanner(s *Scheduler, b banner.Banner) {
	now := time.Now()
	// Convert into seconds
	b.ActiveAt = now.Add(time.Duration(b.ActiveAt) * time.Second).Unix()
	b.ExpireAt = now.Add(time.Duration(b.ExpireAt) * time.Second).Unix()

	fmt.Println("New Banner:", DebugBanner(b, now.Unix()))
	if err := s.Schedule(b); err != nil {
		fmt.Println(err)
	}
	s.Debug(now.Unix())
}
