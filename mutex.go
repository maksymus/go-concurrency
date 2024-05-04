package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var waitingGroup sync.WaitGroup

func main() {
	balance := 100

	//waitingGroup.Add(2000)

	for i := 0; i < 1000; i++ {
		go deposit(&balance, 50)
		go withdraw(&balance, 50)
	}

	waitingGroup.Wait()
	fmt.Println(balance)
}

func deposit(balance *int, amount int) {
	waitingGroup.Add(1)

	mutex.Lock()
	defer mutex.Unlock()

	*balance += amount

	waitingGroup.Done()
}

func withdraw(balance *int, amount int) {
	waitingGroup.Add(1)

	mutex.Lock()
	defer mutex.Unlock()

	*balance -= amount

	waitingGroup.Done()
}
