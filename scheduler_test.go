package banner

import (
	"fmt"
	"testing"
)

func TestScheduler(t *testing.T) {
	fmt.Println("Testing scheduling of banners...")

	// Create a new scheduler
	var s = New()
	var b *Banner
	s.Debug()

	// Schedule banners
	b = &Banner{ID: "A", ActiveAt: 10, ExpireAt: 40}
	fmt.Println("Banner: ", b)
	if err := s.Schedule(b); err != nil {
		fmt.Println(err)
	}
	s.Debug()

	// Schedule banners
	b = &Banner{ID: "B", ActiveAt: 15, ExpireAt: 20}
	fmt.Println("Banner: ", b)
	if err := s.Schedule(b); err != nil {
		fmt.Println(err)
	}
	s.Debug()

}
