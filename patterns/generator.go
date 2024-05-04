package patterns

/**
* Generators return the next value in a sequence each time they are called.
* This means that each value is available as an output before the Generator computes the next value.
*/

func Fibonacci(n int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		i, j := 0, 1
		for k := 0; k < n; k++ {
			ch <- i
			i, j = j, i + j
		}
	}()

	return ch
}
