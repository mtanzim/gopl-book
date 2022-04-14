package memo4

import (
	"fmt"
	"sync"
)

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.Mutex
}

type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Non blocking cache with a mutex to guard the map
// and broadcast channel to signal readiness
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	// first request for key
	// the goroutine becomes responsible for:
	//   - setting the cache value
	//   - broadcasting that value is ready for other goroutines
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		// close the channel to broadcase to others that the value is ready for consumption
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
		fmt.Println("cache hit")
	}
	return e.res.value, e.res.err
}
