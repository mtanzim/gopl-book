package memo6

import (
	"errors"
	"fmt"
)

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Func func(key string, done chan struct{}) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan<- result
	done     chan<- struct{}
}

type Memo struct {
	requests chan request
}

func New(f Func, done chan struct{}) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f, done)
	return memo
}

func (memo *Memo) server(f Func, done chan struct{}) {
	cache := make(map[string]*entry)

	for {
		select {
		case req := <-memo.requests:
			e := cache[req.key]
			// cache miss
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, done)
			}
			go e.deliver(req.response)
		case <-done:
			fmt.Println("cancelled, stopping server")
			// drain channel if needed
			<-memo.requests
			return
		}
	}
}

func (e *entry) deliver(response chan<- result) {
	// wait for cache value to be ready
	<-e.ready
	// return data back to the channel
	response <- e.res
}

func (e *entry) call(f Func, key string, done chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	// broadcase that cache value is ready
	close(e.ready)
}

// Non blocking cache with channels only
func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	select {
	case <-done:
		return nil, errors.New("timer fired, not adding requests for: " + key)
	case memo.requests <- request{key, response, done}:
		res := <-response
		return res.value, res.err
	}

}
