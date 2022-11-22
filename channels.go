package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.

// Create a new channel with make(chan val-type). Channels are typed by the values they convey.

// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.

func main() {

    messages := make(chan string)   // how to make a channel

    go func() { messages <- "ping" }()  // pushing value to a cahnnel

	fmt.Println(messages)  // channel pointer or memory space
	fmt.Println(&messages)  // variable message pointer
	

    msg := <-messages      // pushing stored channel variables to another variable 
    fmt.Println(msg)

}