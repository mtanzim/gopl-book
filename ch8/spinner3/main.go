package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	const n = 45
	fmt.Printf("Calculating Fibonnaci(%d)\n", n)
	go spinner(100 * time.Millisecond)
	go func() {
		res := fib(n)
		fmt.Printf("\rFibonnaci(%d) = %d\n", n, res)
		done <- struct{}{}
	}()
	<-done
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
