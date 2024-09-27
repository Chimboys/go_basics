package main

import (
	"fmt"
	"math/rand"
)

// Chanels in go
// Chanels are a way to share data between goroutines

func basics() {
	//making a channel
	ch := make(chan int)

	//sending data to a channel
	go func() {
		ch <- rand.Intn(100)
	}()

	//receiving data from a channel
	v := <-ch
	fmt.Println(v)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
