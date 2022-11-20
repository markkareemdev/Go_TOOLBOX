package main

import "fmt"

// pointers are just th ememory address of a var
// *int infront of a variable dereferenced the variable from the memory address

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {

    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i) // a pointer address is passed here. the fuction derefrenced it  like this *iptr. an assingned valued to deferenced variable changes the value at that address.
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}