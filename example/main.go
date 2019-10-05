package main

import (
	"net/http"

	"github.com/sangeetk/banner"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "sample.jpg")
}

func Test(t *testing.T) {
	// http.Handle("/", http.HandlerFunc(Handler))
	// log.Println("Running http server at localhost:8080...")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
