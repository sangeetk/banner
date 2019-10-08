package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sangeetk/banner"
)

func main() {

	b := banner.NewBucket("example", datastore.NewMemoryStorage())

	now := time.Now()

	red := &banner.Banner{"red", now.Unix(), now.Add(10 * time.Second).Unix(), "images/red.jpg"}
	b.Put(red)

	http.Handle("/", http.HandlerFunc(Handler))

	log.Println("Running http server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
