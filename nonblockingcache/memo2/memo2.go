package memo

import "sync"

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} //close when res is ready
}

// NonBlockedMemo used broadcasting mechanism, duplicate-suppressing technique
type NonBlockedMemo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

// Func is the types of functions we cache for
type Func func(key string) (interface{}, error)

// New returns the NonBlockedMemo
func New(f Func) *NonBlockedMemo {

	return &NonBlockedMemo{
		f:     f,
		cache: make(map[string]*entry),
	}

}

// Get is the main logic of the non-blocking duplicate-suppressing cache
func (memo *NonBlockedMemo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]

	if e == nil {
		// if not found in cache, this key first seen
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)

	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
