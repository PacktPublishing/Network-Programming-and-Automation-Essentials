package main

import (
    "fmt"
    "runtime"
    "sync"
    "math"
)

var wg sync.WaitGroup

func count() {
    defer wg.Done()
    for i := 1; i < 5; i++ {
        fmt.Println(i)
	var z float64 = 0
	for ; z < 1e8; z++ {
	    _ = math.Sqrt(z)
        }
    }
}

func main() {
    fmt.Println("Version", runtime.Version())
    fmt.Println("NumCPU", runtime.NumCPU())
    fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
    wg.Add(4)
    go count()
    go count()
    go count()
    go count()
    wg.Wait()
}
