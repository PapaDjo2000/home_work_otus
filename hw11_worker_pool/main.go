package main

import (
	"fmt"
	"sync"
)

var count int

func work(a int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
	fmt.Printf("gorutine %d the end \n", a)
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(i, &wg, &mu)
	}
	wg.Wait()
	fmt.Println(count)
}
