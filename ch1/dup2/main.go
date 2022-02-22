package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if (len(files)) == 0 {
		return
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		defer f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts)
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, f.Name())
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}