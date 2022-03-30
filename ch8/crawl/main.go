package main

import (
	"log"
	"os"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

// To avoid spawning a potentially infinite number of goroutines
// resulting in errors such as: 429 Too Many Requests
// The following example shows one can limit the number of concurrent processes
func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string) // deduped urls

	go func() { workList <- os.Args[1:] }()

	// create 20 worker goroutines
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { workList <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
