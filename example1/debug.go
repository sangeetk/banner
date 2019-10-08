package main

import (
	"fmt"

	"github.com/sangeetk/banner"
)

// DebugBanner is used for testing purpose
func DebugBanner(b *banner.Banner, offset int64) string {
	if b == nil {
		return " [null,null)"
	}
	return fmt.Sprintf("%v[%v,%v)", b.ID, b.ActiveAt-offset, b.ExpireAt-offset)
}

// DebugScheduler displays schedule on console
func DebugScheduler(s *banner.Scheduler, offset int64) {
	// Display linked list
	fmt.Print("Head -> ")
	for t := s.Head; t != nil; t = t.Next {
		// for i := t.T1; i < t.T2; i++ {
		//	fmt.Print(t.Banners[0].ID)
		// }
		fmt.Printf("{'%v', %v-%v}", t.Banners[0].ID, t.T1-offset, t.T2-offset)
		fmt.Print(" -> ")
	}
	fmt.Println("nil")

	// Display timeline
	var start, end int64
	if s.Head != nil {
		start = s.Head.T1
	}
	fmt.Printf("Timeline: (%02d)", start-offset)
	for t := s.Head; t != nil; t = t.Next {
		end = t.T2
	}
	for i := start; i <= end; i++ {
		if i%10 == 0 {
			fmt.Print("|")
		} else if i%5 == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Printf("(%02d)\n", end-offset)

	// Display scheduled banners
	fmt.Printf("Scheduling:   ")
	for i := start; i <= end; i++ {
		b, err := s.GetByTime(i)
		if err != nil {
			fmt.Print(" ")
		} else {
			fmt.Print(b.ID)
		}
	}
	fmt.Println()
	fmt.Println()
}
