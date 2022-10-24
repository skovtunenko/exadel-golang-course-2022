package mapsexample

import (
	"sync"
	"testing"
)

func TestConcurrentAccess(t *testing.T) {
	const N = 10

	m := make(map[int]int)

	wg := &sync.WaitGroup{}

	for i := 0; i < N; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			m[i] = i * 10
		}(i)
	}

	wg.Wait()
	t.Log(m)
}
