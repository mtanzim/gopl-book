package main

import (
	"fmt"
	"time"
)

func main() {
	// `go` here allows the subsequent LOC to execute while the spinner is still operating
	go spinner(100 * time.Millisecond)
	const n = 45
	fmt.Printf("Calculating Fibonnaci(%d)\n", n)
	fibN := fib(n)
	fmt.Printf("\rFibonnaci(%d) = %d\n", n, fibN)
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
