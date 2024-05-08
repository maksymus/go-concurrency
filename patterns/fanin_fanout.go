package patterns

import "sync"

/**
* Fan-In refers to a technique in which you join data from multiple inputs into a single entity.
* On the other hand, Fan-Out means to divide the data from a single source into multiple smaller chunks.
*/


func FanIn(chs ... chan string) <-chan string {
	result := make(chan string)

	go func() {
		var wg sync.WaitGroup

		defer close(result)
		for _, ch := range chs {
			wg.Add(1)
			go func(channel chan string) {
				defer wg.Done()
				for str := range channel {
					result <- str
				}
			}(ch)
		}

		wg.Wait()
	}()

	return result
}
