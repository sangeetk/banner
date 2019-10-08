package main

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	// Get active banner
	b, err := Coll.GetActive()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/jpg")
	w.Write(b.Image)

}
