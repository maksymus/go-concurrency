package main

import "context"

func init() {
	JobQueue = make(chan Job, 100) // Buffered channel for job queue
}

var JobQueue chan Job

type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(numWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, numWorkers)
	return &Dispatcher{
		workerPool: workerPool,
		maxWorkers: numWorkers,
	}
}

func (d *Dispatcher) Run(ctx context.Context) {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.workerPool)
		worker.Start(ctx)
	}

	go d.dispatch(ctx)
}

func (d *Dispatcher) dispatch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// if context is done, stop the dispatcher
			return
		case job := <-JobQueue:
			go func(job Job) {
				// get a worker from the pool
				jobChannel := <-d.workerPool

				// send the job to the worker
				jobChannel <- job
			}(job)
		}
	}
}
