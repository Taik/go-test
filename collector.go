package main

import (
	"fmt"
	"net/http"
	"time"
)

// A buffered channel that we can send requests on.
var WorkQueue = make(chan WorkRequest, 100)

func Collector(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse out the delay from `delay` field
	delay, err := time.ParseDuration(r.PostFormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate delay
	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "The delay must be between 1-10 seconds, inclusively.", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if len(name) == 0 {
		http.Error(w, "The task must have a name.", http.StatusBadRequest)
		return
	}

	work := WorkRequest{Name: name, Delay: delay}

	WorkQueue <- work

	fmt.Println("New task queued.")

	w.WriteHeader(http.StatusCreated)
	return
}
