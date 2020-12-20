package memo

import "sync"

// BlockedMemo uses double mutex
type BlockedMemo struct {
	f     Func
	cache map[string]result
	mutex sync.Mutex
}

// Func is the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New returns a new blocked cache
func New(f Func) *BlockedMemo {
	return &BlockedMemo{f: f, cache: make(map[string]result)}
}

// Get is the main logic of content retrieving with cache
func (memo *BlockedMemo) Get(key string) (interface{}, error) {
	memo.mutex.Lock()
	res, ok := memo.cache[key]
	memo.mutex.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mutex.Lock()
		memo.cache[key] = res
		memo.mutex.Unlock()
	}
	return res.value, res.err
}
