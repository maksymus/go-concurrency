package main

import "fmt"

type Payload struct{}

func (p Payload) Process() error {
	fmt.Println("Payload processing")
	return nil
}

type Job struct {
	payload Payload
}

type Worker struct {
	workerPool chan chan Job
	jobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		workerPool: workerPool,
		jobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			// add worker to the pool
			w.workerPool <- w.jobChannel

			select {
			case job := <-w.jobChannel:
				// process the job
				if err := job.payload.Process(); err != nil {
					fmt.Println("Error processing job:", err)
				}
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- true
		close(w.jobChannel)
		close(w.quit)
	}()
}
