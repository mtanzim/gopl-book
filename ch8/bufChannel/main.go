package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func request(url string) string {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func query() string {
	responses := make(chan string, 3)
	go func() { responses <- request("https://jsonplaceholder.typicode.com/posts/1") }()
	go func() { responses <- request("https://jsonplaceholder.typicode.com/posts/2") }()
	go func() { responses <- request("https://jsonplaceholder.typicode.com/posts/3") }()
	return <-responses
}

func leakyQuery() string {
	responses := make(chan string)
	go func() { responses <- request("https://jsonplaceholder.typicode.com/posts/1") }()
	go func() { responses <- request("https://jsonplaceholder.typicode.com/posts/2") }()
	go func() { responses <- request("https://jsonplaceholder.typicode.com/posts/3") }()
	return <-responses
}

func main() {
	fmt.Println(query())
	fmt.Println(leakyQuery())
}
