package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
}

func main() {
	personPool := sync.Pool{
		New: func() any {
			return new(Person)
		},
	}

	newPerson := personPool.Get().(*Person)
	newPerson.name = "Max"
	personPool.Put(newPerson)

	newPerson1 := personPool.Get().(*Person)
	fmt.Println(newPerson1.name)
}
