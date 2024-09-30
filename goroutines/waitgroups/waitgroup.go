package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// WaitGroup is a struct that allows you to wait for a collection of goroutines to finish
// It is a counter that is incremented by the number of goroutines you want to wait for
// Each goroutine decrements the counter when it finishes
// When the counter reaches 0, the Wait() method unblocks
// Done() decrements the counter, and Add() increments it
// Defer works well with WaitGroup because it ensures that the counter is decremented even if the goroutine panics
// Defer can be used with any function, not just with WaitGroup
// Defer can help to treat panic as an error

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done() //Defer is used to ensure that the counter is decremented even if the goroutine panics
		for i := 0; i < 20; i++ {
			fmt.Println("First goroutine:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			// Defer can help to treat panic as an error
			// Recover is built-in function that stops the panic and returns the value that was passed to the call to panic
			// Recover is only useful inside deferred functions !!!
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()

		for i := 0; i < 20; i++ {
			fmt.Println("Second goroutine:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	fmt.Println("To show that it does finish one routine before the other")

	wg.Wait() // without wait the main goroutine will finish before the other goroutines,
	// so the other goroutines will not finish
	fmt.Println("All goroutines finished")

	// WaitGroup with multiple goroutines
	// Example 2
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
	}

	// Launch a goroutine for each URL
	for _, url := range urls {
		wg.Add(1) // Increment the WaitGroup counter
		go fetchURL(url, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All requests completed.")

}

func fetchURL(url string, wg *sync.WaitGroup) {
	// Notify the WaitGroup that we're done with this goroutine
	defer wg.Done()

	// Send a GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Printf("Fetched %s with status: %s\n", url, resp.Status)
}
