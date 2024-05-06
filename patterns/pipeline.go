package patterns

/**
* A pipeline is a series of independent components or stages connected via connectors.
* The component obtains data from the previous stage, performs some operations, and emits the data onto the next stage.
*/

func Generate(nums ...int) chan int {
    ch := make(chan int)

    go func() {
        defer close(ch)
        for _, num := range nums {
            ch <- num
        }
    }()

    return ch
}

func Modify(data chan int, op func(int) int) chan int {
    ch := make(chan int)

    go func() {
        defer close(ch)
        for num := range data {
            ch <- op(num)
        }
    }()

    return ch
}
