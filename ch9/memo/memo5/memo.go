package memo5

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
	cache := make(map[string]result)
	for req := range memo.requests {
		res, ok := cache[req.key]
		if ok {
			req.response <- res
		} else {
			res := result{}
			res.value, res.err = f(req.key)
			cache[req.key] = res
			req.response <- res
		}
	}
}

// Non blocking cache with channels
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}
