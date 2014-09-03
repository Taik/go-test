package main

import (
	"fmt"
	"time"
)

func main() {
	// Initialize a new worker pool & start workers
	pool := WorkerPool{}
	pool.Start(4)
	time.Sleep(time.Second * 1)

	// Broadcasting stop to all workers
	fmt.Println("Stopping..")
	pool.Stop()
	fmt.Println("Stopped")

	// Wait for goroutines to actually stop
	time.Sleep(time.Second * 2)
}
