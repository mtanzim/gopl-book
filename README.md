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
  - [spinner](ch8/spinner/main.go) 
  - [clock](ch8/clock2/main.go)
  - [echo server](ch8/reverb2/main.go)

#### Channels

- Goroutines are the activities of concurrent programs, and `channels` are the connections between them
- It is a communication mechanism through which one goroutine can can send values to another goroutine
- Channels carry value of an _element type_
- The built in `make` function can create channels

```go
ch := make (chan int)
```

- Channels are reference types similar to maps and slices. Therefore, copying channels, passing them as arguments copies a _reference_ referring to the same data structure
- The zero value of channels is `nil`
- Channels allow two operations, _send_ and _receive_, collectively known as _communications_; both use the `<-` operator

```go
ch <- x // a send statement
x = <- ch // a receive statement
<- ch // receive, discard result
```
- Channels also support a `close` operation
- Closed channels indicate that no more values will be sent; subsequent attempts at send will panic
- Closed channels can be received from until drained, and all values after will be the zero value of the channel _element type_
- Channels can be _buffered_ or _unbuffered_, _unbuffered_ channels have non-zero _capacity_; details will be explained in the following section

```go
ch = make(chan int) // unbuffered channel
ch = make(chan int) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity of 3
```

##### Unbuffered Channels






