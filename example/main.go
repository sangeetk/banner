package main

import (
	"fmt"
	"time"

	"github.com/sangeetk/banner"
)

func main() {
	fmt.Println("Testing scheduling of banners...")

	// Create a new collection
	var col = banner.NewCollection("collection1", banner.NewDataStoreMemory())

	now := time.Now()

	// Time here is in seconds from current time
	red, _ := banner.NewFile("R", "images/red.jpg")
	red.AddDuration(now.Add(0*time.Second), now.Add(20*time.Second))
	col.Add(*red)
	debugScheduler(&col.Scheduler, now.Unix())

	blue, _ := banner.NewFile("B", "images/blue.jpg")
	blue.AddDuration(now.Add(40*time.Second), now.Add(60*time.Second))
	col.Add(*blue)
	debugScheduler(&col.Scheduler, now.Unix())

	green, _ := banner.NewFile("G", "images/green.jpg")
	green.AddDuration(now.Add(25*time.Second), now.Add(45*time.Second))
	col.Add(*green)
	debugScheduler(&col.Scheduler, now.Unix())

	yellow, _ := banner.NewFile("Y", "images/yellow.jpg")
	yellow.AddDuration(now.Add(70*time.Second), now.Add(100*time.Second))
	col.Add(*yellow)
	debugScheduler(&col.Scheduler, now.Unix())

	// Remove green
	col.Remove(*green)
	debugScheduler(&col.Scheduler, now.Unix())

	orange, _ := banner.NewFile("O", "images/orange.jpg")
	orange.AddDuration(now.Add(90*time.Second), now.Add(120*time.Second))
	col.Add(*orange)
	debugScheduler(&col.Scheduler, now.Unix())

	green, _ = banner.NewFile("G", "images/green.jpg")
	green.AddDuration(now.Add(35*time.Second), now.Add(65*time.Second))
	col.Add(*green)
	debugScheduler(&col.Scheduler, now.Unix())

	pink, _ := banner.NewFile("P", "images/pink.jpg")
	pink.AddDuration(now.Add(120*time.Second), now.Add(130*time.Second))
	col.Add(*pink)
	debugScheduler(&col.Scheduler, now.Unix())

	// default banner
	white, _ := banner.NewFile(".", "images/white.jpg")
	white.AddDuration(now.Add(0*time.Second), now.Add(160*time.Second))
	col.Add(*white)
	debugScheduler(&col.Scheduler, now.Unix())

}
