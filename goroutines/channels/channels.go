package main

// go run goroutines/channels/channels.go
// go run goroutines/channels/channels.go goroutines/channels/basics_channels.go

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0, 12, 10, 22, 28, 231}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from chanel

	fmt.Println(x, y, x+y)

	BufferedChannels()
	Select()
}

//The order in which x and y are assigned is not guaranteed to match the order in which the goroutines are started.
// Instead, it reflects the order in which the computations are completed.
// This is a typical use case for channels, allowing you to synchronize and collect results from concurrent operations.
