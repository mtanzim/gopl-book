package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	var done = make(chan struct{})
	var wg sync.WaitGroup
	cancelled := func() bool {
		select {
		// once done is closed, this channel will receive zero values
		case <-done:
			return true
		default:
			return false
		}
	}

	// goroutine to listen for input on STDIN
	wg.Add(1)
	go func() {
		defer wg.Done()
		os.Stdin.Read(make([]byte, 1))
		// trigger broadcast by closing channel
		close(done)
	}()

	// set up ticker functions
	tickerFn := func(n int, s string) {
		defer wg.Done()
		tick := time.Tick(time.Duration(n) * time.Second)
		for {
			if cancelled() {
				fmt.Printf("=== %s cancelled ===\n", s)
				return
			}
			fmt.Println(s)
			<-tick
		}

	}

	// kick off ticker functions
	wg.Add(4)
	go tickerFn(1, "Hello")
	go tickerFn(2, "Numa Numa")
	go tickerFn(3, "World")
	go tickerFn(5, "Joomla")
	// wait for all goroutines to return
	wg.Wait()

}
