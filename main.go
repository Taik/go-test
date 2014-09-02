package main

import (
	"fmt"
	"time"
)

func main() {
	pool := WorkerPool{}
	pool.Start(4)
	time.Sleep(time.Second * 1)
	fmt.Println("Stopping..")
	pool.Stop()
	fmt.Println("Stopped")
	time.Sleep(time.Second * 2)
}
