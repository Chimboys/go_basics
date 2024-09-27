package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// WaitGroup is a struct that allows you to wait for a collection of goroutines to finish
// It is a counter that is incremented by the number of goroutines you want to wait for
// Each goroutine decrements the counter when it finishes
// When the counter reaches 0, the Wait() method unblocks
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			fmt.Println("First goroutine:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			fmt.Println("Second goroutine:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	fmt.Println("To show that it does finish one routine before the other")

	wg.Wait()
	fmt.Println("All goroutines finished")

}
