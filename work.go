package main

import "time"

// Some unit of work
type WorkRequest struct {
	Name  string
	Delay time.Duration
}
