package memo5

import "fmt"

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		// cache miss
		if e == nil {
			fmt.Println("cache miss for: " + req.key)
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) deliver(response chan<- result) {
	// wait for cache value to be ready
	<-e.ready
	// return data back to the channel
	response <- e.res
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	// broadcase that cache value is ready
	close(e.ready)
}

// Non blocking cache with channels only
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}
