package patterns

import (
    "sync"
)

func Work(nums []int, numWorkers int, f func(int) int) (ret []int) {
    in := make(chan int, numWorkers)
    out := make(chan int, numWorkers)

    var wg sync.WaitGroup

    // start workers
    go func() {
        defer close(out)
        defer wg.Wait()
        for i := 0; i < numWorkers; i++ {
            wg.Add(1)
            go work(in, out, &wg, f)
        }
    }()

    // generate data
    go func() {
        defer close(in)
        for _, num := range nums {
            in <- num
        }
    }()

    // process result
    for res := range out {
       ret = append(ret, res)
    }

    return ret
}

func work(in <-chan int, out chan<- int, wg *sync.WaitGroup, f func(int) int) {
    defer wg.Done()

    for num := range in {
        out <- f(num)
    }
}
