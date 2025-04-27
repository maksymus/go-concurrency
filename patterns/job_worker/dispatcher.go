package main

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

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.workerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
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
