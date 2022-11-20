// A goroutine is a lightweight thread of execution.
// since thread execution is concurent in go by default. goroutine is a way to run things synchronously

// To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one
// Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).

package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}