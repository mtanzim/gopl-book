package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	// memo "github.com/mtanzim/gopl-book/ch9/memo/memo1"
	// memo "github.com/mtanzim/gopl-book/ch9/memo/memo2"
	// memo "github.com/mtanzim/gopl-book/ch9/memo/memo3"
	// memo "github.com/mtanzim/gopl-book/ch9/memo/memo4"
	// memo "github.com/mtanzim/gopl-book/ch9/memo/memo5"
	memo "github.com/mtanzim/gopl-book/ch9/memo/memo6"
)

func httpGetBody(url string, done chan struct{}) (interface{}, error) {

	cancelled := func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	}

	if cancelled() {
		return nil, errors.New("Execution cancelled")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingUrls() []string {
	return []string{
		"https://golang.org",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://gopl.io",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://gopl.io",
	}
}

func TestMainSequential(t *testing.T) {
	m := memo.New(httpGetBody, nil)
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url, nil)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s\t\t %s\t\t %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func TestMainConcurrent(t *testing.T) {
	m := memo.New(httpGetBody, nil)
	var n sync.WaitGroup
	for _, url := range incomingUrls() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, nil)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s\t\t %s\t\t %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func TestMainConcurrentWithCancel(t *testing.T) {
	done := make(chan struct{})
	m := memo.New(httpGetBody, done)
	var n sync.WaitGroup
	for _, url := range incomingUrls() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s\t\t %s\t\t %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}
