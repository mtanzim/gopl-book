# gopl-book

[Book](http://www.gopl.io/)

## Why?

- Refresher on Golang basics
- Improve understanding of Go concurrency mechanisms

## Related efforts

- [UCI Go MOOC Exercises](https://github.com/mtanzim/go-uci)

## Notes on critical concepts

### Interfaces

- TODO

### Goroutines and Channels

- Go allows for two styles of concurrent programming
  1. __Communicating sequential processes (CSP)__: passing values between independent activities (goroutines)
  1. __Shared memory multithreading__ (threads in most mainstream languages)

#### Goroutines

- Each concurrently executing activity is called a goroutine
- For example, imagine 2 independent functions
- In a sequential program, the functions are called one after the other
- In a concurrent program, both functions are active at once
- Goroutines are similar to an OS thread. This is a fair assumption for writing correct programs. Differences will be described later
- Upon a program start, the only goroutine is the `main` function; this is known as the `main goroutine`
- New goroutines can be created with the `go` statement, ie:

```go
f() // call f, await return
go f() // create a goroutine to call f(), DO NOT wait
```

- Examples
  - [spinner](ch8/spinner/main.go) for a more elaborate example
  - [clock](ch8/clock2/main.go) for a more elaborate example
  - [echo server](ch8/reverb2/main.go) for a more elaborate example

