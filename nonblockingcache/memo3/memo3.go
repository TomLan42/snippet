package memo

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} //close when res is ready
}

type request struct {
	key      string
	response chan<- result //This result channel is write-only, so it is chan<-
}

// MonitoredMemo uses monitor goroutine to manage things
type MonitoredMemo struct {
	requests chan request
}

// Func is the types of functions we cache for
type Func func(key string) (interface{}, error)

// New returns the MonitoredMemo
func New(f Func) *MonitoredMemo {
	memo := &MonitoredMemo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Get is the main logic of the non-blocking duplicate-suppressing cache
func (memo *MonitoredMemo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

// Close the requests channel of MonitoredMemo
func (memo *MonitoredMemo) Close() {
	close(memo.requests)
}

//This part is called the monitor goroutine

func (memo *MonitoredMemo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		//cache miss
		if e == nil {
			e = &entry{
				ready: make(chan struct{}),
			}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res

}
