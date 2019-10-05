package main

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/sample.jpg")
}
