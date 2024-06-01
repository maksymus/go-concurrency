package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    go func() {
        for i := 0; i < 5; i++ {
            ch <- fmt.Sprintf("chan1 message each 1 second %d\n", i)
            time.Sleep(time.Second)
        }
    }()

    timer := time.After(3 * time.Second)

    for {
        select {
        case message1 := <-ch:
            fmt.Printf(message1)
        case <-timer:
            fmt.Println("3 seconds passed. Exiting...")
            return
        }
    }
}
