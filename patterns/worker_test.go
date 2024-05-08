package patterns

import (
    "runtime"
    "testing"
    "time"
)

var ints = generateInt(1000)

func BenchmarkWork1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Work(ints, 1, pow2)
    }
}

func BenchmarkWork2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Work(ints, 2, pow2)
    }
}

func BenchmarkWork5(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Work(ints, 5, pow2)
    }
}

func BenchmarkWork10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Work(ints, 10, pow2)
    }
}

func BenchmarkWork100(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Work(ints, 100, pow2)
    }
}

func BenchmarkWork1000(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Work(ints, 1000, pow2)
    }
}

func BenchmarkWorkNCPU(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Work(ints, runtime.NumCPU(), pow2)
    }
}

func generateInt(count int) []int {
    res := make([]int, 0)

    for i := 1; i <= count; i++ {
        res = append(res, i)
    }

    return res
}

func pow2(n int) int {
    time.Sleep(time.Millisecond * 2)
    return n * n
}
