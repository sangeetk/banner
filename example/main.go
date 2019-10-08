package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sangeetk/banner"
)

var Coll *banner.Collection
var startTime time.Time

func main() {
	fmt.Println("Testing scheduling of banners...")

	// Create a new collection
	Coll = banner.NewCollection("collection1", banner.NewDataStoreMemory())

	startTime = time.Now()

	// Time here is in seconds from current time
	red, _ := banner.NewFile("R", "images/red.jpg")
	red.AddDuration(startTime.Add(0*time.Second), startTime.Add(20*time.Second))
	Coll.Add(*red)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	blue, _ := banner.NewFile("B", "images/blue.jpg")
	blue.AddDuration(startTime.Add(40*time.Second), startTime.Add(60*time.Second))
	Coll.Add(*blue)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	green, _ := banner.NewFile("G", "images/green.jpg")
	green.AddDuration(startTime.Add(25*time.Second), startTime.Add(45*time.Second))
	Coll.Add(*green)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	yellow, _ := banner.NewFile("Y", "images/yellow.jpg")
	yellow.AddDuration(startTime.Add(70*time.Second), startTime.Add(100*time.Second))
	Coll.Add(*yellow)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	// Remove green
	Coll.Remove(*green)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	orange, _ := banner.NewFile("O", "images/orange.jpg")
	orange.AddDuration(startTime.Add(90*time.Second), startTime.Add(120*time.Second))
	Coll.Add(*orange)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	green, _ = banner.NewFile("G", "images/green.jpg")
	green.AddDuration(startTime.Add(35*time.Second), startTime.Add(65*time.Second))
	Coll.Add(*green)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	pink, _ := banner.NewFile("P", "images/pink.jpg")
	pink.AddDuration(startTime.Add(120*time.Second), startTime.Add(130*time.Second))
	Coll.Add(*pink)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	mercari, _ := banner.NewFile("M", "images/mercari.jpg")
	mercari.AddDuration(startTime.Add(0*time.Second), startTime.Add(10*time.Second))
	Coll.Add(*mercari)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	mercari2, _ := banner.NewURL("m", "https://web-jp-assets.mercdn.net/_next/static/images/top-banner-super-exhibition-fes-1354ceda34bd06081a45ee755e911f07.jpg")
	mercari2.AddDuration(startTime.Add(160*time.Second), startTime.Add(180*time.Second))
	Coll.Add(*mercari2)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	// default banner
	white, _ := banner.NewFile(".", "images/white.jpg")
	white.AddDuration(startTime.Add(0*time.Second), startTime.Add(160*time.Second))
	Coll.Add(*white)
	debugScheduler(&Coll.Scheduler, startTime.Unix())

	// Start http server
	go func() {
		http.Handle("/", http.HandlerFunc(Handler))
		log.Println("Running http server at localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Ask for inputs from user
	time.Sleep(1 * time.Second)

	for {
		var sec int64
		fmt.Print("\nPreview time in sec : ")
		fmt.Scanf("%v", &sec)

		var active, preview string
		if b, err := Coll.GetActive(); err != nil {
			active = "none"
		} else {
			active = b.ID
		}

		t := startTime.Add(time.Duration(sec) * time.Second)
		if b, err := Coll.Preview(t); err != nil {
			preview = "none"
		} else {
			preview = b.ID
		}

		fmt.Printf("Time [%v sec] => Active startTime : [%v] , Active at %v sec : [%v]\n", time.Now().Unix()-startTime.Unix(), active, sec, preview)
	}

}
