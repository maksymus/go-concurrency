package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			chan1 <- fmt.Sprintf("chan1 message each 1 second %d\n", i)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			chan2 <- fmt.Sprintf("chan2 message each 500 ms %d\n", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case message1 := <-chan1:
			fmt.Printf(message1)
		case message2 := <-chan2:
			fmt.Printf(message2)
		}
	}
}
