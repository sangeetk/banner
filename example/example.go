package main

import (
	"log"
	"net/http"

	"github.com/sangeetk/banner"
	"github.com/sangeetk/banner/datastore"
)

func main() {

	b := banner.Init(datastore.InMemory())

	b.Put()
	b.Put()

	b.Get()

	b.Del()

	http.Handle("/", http.HandlerFunc(Handler))

	log.Println("Running http server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
