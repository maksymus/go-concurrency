package main

import (
	"fmt"
)

func main() {
	ints := []int { 1, -5, 3, 7, 6, 0, -2 }

	ch := make(chan int)

	go gosum(ch, ints[:len(ints)/2])
	go gosum(ch, ints[len(ints)/2:])

	x, y := <- ch, <- ch

	fmt.Println(ints[len(ints)/2:])
	fmt.Println(x, y, x + y)

	// send channel
	comm := make(chan int)
	go send(comm)

	for i := 0; i < 5; i++ {
		fmt.Println(<-comm)
	}

	close(comm)
}

func gosum(ch chan int, ints []int)  {
	ch <- sum(ints)
}

func sum(ints []int) int {
	var sum int
	for _, i := range ints {
		sum += i
	}

	return sum
}

func send(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}
