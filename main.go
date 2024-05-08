package main

import (
	"fmt"
	"github.com/maksymus/go-concurrency/patterns"
	"strconv"
	"time"
)

func main() {

	// Generator pattern
	generator := patterns.Fibonacci(10)
	for i := range generator {
		fmt.Println(i)
	}

	// Fan-In, Fan-Out
	ch1 := channelGenerator(1, 2, 3, 4, 5, 6)
	ch2 := channelGenerator(120, 230, 600, 800)

	fanIn := patterns.FanIn(ch1, ch2)
	for num := range fanIn {
		fmt.Println("fanIn: " + num)
	}

	// Semaphore
	semaphore := make(patterns.Semaphore, 3)
	for i := 0; i < 10; i++ {
		go func(s patterns.Semaphore, i int) {
			s.Acquire()
			defer s.Release()

			fmt.Printf("running thread: %d\n", i)
			time.Sleep(time.Second)
		}(semaphore, i)
	}

	// Pipeline
	ch := patterns.Modify(patterns.Generate(1, 2, 3, 4), pow2)
	for num := range ch {
		fmt.Printf("pipeline: %d\n", num)
	}

	// Worker Pool
	res := patterns.Work([]int{1, 2, 3, 4, 5, 6, 7}, 3, pow2)
	fmt.Printf("workers: %v\n", res)
}

func channelGenerator(numbers ... int) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, num := range numbers{
			//fmt.Println("num: " + strconv.Itoa(num))
			ch <- strconv.Itoa(num)
		}
	}()
	return ch
}

func pow2(n int) int {
	return n * n
}
