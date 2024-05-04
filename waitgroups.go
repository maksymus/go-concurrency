package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("thread 1")
		time.Sleep(time.Millisecond * 2)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("thread 2")
		time.Sleep(time.Millisecond * 100)
	}()

	wg.Wait()
}
