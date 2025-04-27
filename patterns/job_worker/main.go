package main

import "time"

func main() {
	dispatcher := NewDispatcher(3)
	dispatcher.Run()

	JobQueue = make(chan Job, 10)

	for i := 0; i < 10; i++ {
		job := Job{payload: Payload{}}
		JobQueue <- job
	}

	// Wait for all jobs to be processed
	// This is just a placeholder. In a real application, you would use a sync.WaitGroup or similar mechanism.
	time.Sleep(time.Second * 5)
}
