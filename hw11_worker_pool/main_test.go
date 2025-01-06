package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var count1 int

func TestWork(t *testing.T) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(i, &wg, &mu)
		wg.Wait()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				mu.Lock()
				count1++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	assert.Equal(t, count1, count)
}
