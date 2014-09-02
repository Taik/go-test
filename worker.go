package main

import (
	"fmt"
	"time"
)

type Worker struct {
	ID       int
	WorkChan chan WorkRequest
	QuitChan chan bool
}

func (w *Worker) Start() {
	// Goroutine!
	go func() {
		select {
		case work := <-w.WorkChan:
			fmt.Printf("Worker #%d: Received task, delaying for %f seconds.\n", w.ID, work.Delay)
			time.Sleep(work.Delay)
			fmt.Printf("Worker #%d: Hello %s!\n", w.ID, work.Name)

		case <-w.QuitChan:
			fmt.Printf("Worker #%d: Stopping worker.\n", w.ID)
			return
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func NewWorker(id int, workQueue chan WorkRequest) *Worker {
	worker := &Worker{
		ID:       id,
		WorkChan: workQueue,
		QuitChan: make(chan bool),
	}
	return worker
}

// Worker slice to hold all current workers
type WorkerPool struct {
	Workers []*Worker // By default, this contains 0 workers
}

func (w *WorkerPool) Start(numWorkers int) {
	var current_worker *Worker
	for i := 0; i < numWorkers; i++ {
		current_worker = NewWorker(i+1, WorkQueue)
		current_worker.Start()
		w.Workers = append(w.Workers, current_worker)
		fmt.Printf("Worker #%d starting up! (@ %p)\n", current_worker.ID, current_worker)
	}
}

func (w *WorkerPool) Stop() {
	for _, current_worker_ptr := range w.Workers {
		current_worker := *current_worker_ptr
		fmt.Printf("Stopping worker #%d (@ %p)\n", current_worker.ID, &current_worker)
		current_worker.Stop()
	}
}
