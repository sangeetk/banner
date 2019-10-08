package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/sangeetk/banner"
)

func main() {
	fmt.Println("Testing scheduling of banners...")

	// Create a new collection
	var col = banner.NewCollection("collection1", banner.NewDataStoreMemory())

	// Time here is in seconds from current time
	b1 := newBanner(&banner.Banner{ID: "A", ActiveAt: 10, ExpireAt: 80})
	col.Add(b1)
	DebugScheduler(col.Scheduler)

	b2 := newBanner(&banner.Banner{ID: "B", ActiveAt: 20, ExpireAt: 70})
	col.Add(b2)
	DebugScheduler(col.Scheduler)

	b3 := newBanner(&banner.Banner{ID: "C", ActiveAt: 30, ExpireAt: 60})
	col.Add(b3)
	DebugScheduler(col.Scheduler)

	b4 := newBanner(&banner.Banner{ID: "D", ActiveAt: 40, ExpireAt: 50})
	col.Add(b4)
	DebugScheduler(col.Scheduler)

	// default banner
	dflt := newBanner(&banner.Banner{ID: ".", ActiveAt: 0, ExpireAt: 100})
	col.Add(dflt)
	DebugScheduler(col.Scheduler)

}

func newBanner(s *Scheduler, b *banner.Banner) {
	now := time.Now()
	b.ActiveAt = now.Add(time.Duration(b.ActiveAt) * time.Second).Unix()
	b.ExpireAt = now.Add(time.Duration(b.ExpireAt) * time.Second).Unix()
}
