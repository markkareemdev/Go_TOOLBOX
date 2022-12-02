package main

import "fmt"

func ping(pings chan<- string, msg string) { 
	// note the direction on the function parameter it specifies the channel ping should only receive here
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string, ponger chan<- string) {
    //  note 3 channels here but does different things.
	//  pings channel sends while pongs and ponger receives
	//  only channels can receive from a channel
	msg := <-pings

	fmt.Println(msg,`i am inside pong function`)
    pongs <- msg
	ponger <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
	ponger := make(chan string, 1)

    ping(pings, "passed message")
    pong(pings, pongs, ponger)

    fmt.Println(<-pongs)
	fmt.Println(<-ponger)
}