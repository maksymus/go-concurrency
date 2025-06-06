package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up channel to listen for interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle interrupt signals
	go func() {
		<-signalChan
		cancel()
	}()

	// Example usage of the dispatcher and job queue
	dispatcher := NewDispatcher(3)
	dispatcher.Run(ctx)

	for i := 0; i < 10; i++ {
		job := Job{payload: Payload{}}
		JobQueue <- job
	}

	// Wait for all jobs to be processed
	// This is just a placeholder. In a real application, you would use a sync.WaitGroup or similar mechanism.
	time.Sleep(time.Second * 5)
}
