package main

import (
	"fmt"
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
// The following example shows how one can limit the number of concurrent processes
type linkContainer struct {
	link  string
	depth int
}

func main() {
	workList := make(chan []*linkContainer)
	unseenLinks := make(chan *linkContainer) // deduped urls

	go func() {
		lst := []*linkContainer{}
		for _, l := range os.Args[1:] {
			lst = append(lst, &linkContainer{l, 0})
		}
		workList <- lst
	}()

	// create 20 worker goroutines, limit concurrency
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.link)
				lst := []*linkContainer{}
				for _, l := range foundLinks {
					lst = append(lst, &linkContainer{l, link.depth + 1})
				}
				go func() { workList <- lst }()
			}
		}()
	}

	seen := make(map[string]bool)

	for list := range workList {
		for _, link := range list {
			if !seen[link.link] {
				if link.depth > 1 {
					fmt.Printf("%s exceeds depth requirements as it is of depth %d\n", link.link, link.depth)
					continue
				}
				seen[link.link] = true
				unseenLinks <- link
			}
		}
	}
}
