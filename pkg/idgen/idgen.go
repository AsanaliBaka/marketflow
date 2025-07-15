package idgen

import "sync"

func NewIdGeneretor() func() int64 {
	var mu sync.Mutex
	var counter int64 = 0

	return func() int64 {
		mu.Lock()
		defer mu.Unlock()

		counter++
		return counter
	}
}
