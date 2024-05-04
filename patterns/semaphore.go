package patterns

type Semaphore chan any

func (s Semaphore) Acquire() {
	s <- struct {}{}
}

func (s Semaphore) Release() {
	<-s
}

