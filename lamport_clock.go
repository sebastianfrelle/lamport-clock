package main

import (
	"fmt"
	"time"
)

// Buffer capacity for interprocess channels
const chanBufferCap = 1000

func main() {
	// Interprocess channels
	c12 := make(chan int, chanBufferCap)
	c21 := make(chan int, chanBufferCap)

	// Processes
	p1 := Process{0, 0, 1 * time.Second, c21, c12}
	p2 := Process{1, 0, 5 * time.Second, c12, c21}

	fmt.Println("starting p1")
	p1.Start()
	fmt.Println("starting p2")
	p2.Start()

	// Block main thread forever
	<-make(chan bool)
}
