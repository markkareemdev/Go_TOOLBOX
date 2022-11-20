package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))

    var fib func(n int) int

    fib = func(n int) int {
		var count = 0
        if n < 2 {
			count += 1
			fmt.Println("count: ",count)
            return n
        }

        return fib(n-1) + fib(n-2)
    }

    fmt.Println(fib(4))
}