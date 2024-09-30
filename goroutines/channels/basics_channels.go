package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Chanels in go
// Chanels are a way to share data between goroutines

func Basics() {
	//making a channel
	ch := make(chan int)

	//sending data to a channel
	//started a goroutine with go
	go func() {
		ch <- rand.Intn(100)
	}()

	//receiving data from a channel
	v := <-ch
	fmt.Println(v)
}

// initializing a channel with a buffer (capacity)
func BufferedChannels() {
	channel := make(chan int, 3) //buffered channel with capacity 2
	channel <- 1
	channel <- 2
	channel <- 3
	// channel <- 4 //this will cause a deadlock

	for i := 0; i < 3; i++ {
		fmt.Println(<-channel)
	}

}

//select statement

func Select() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from Channel 1, every second"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c2 <- "from Channel 2, every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	//without select, c1 will print, then it will 2 seconds for channel 2 to print, and again c1 will print
	// with select, it will not wait for the other channel to print, it will print the channel that is ready
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
