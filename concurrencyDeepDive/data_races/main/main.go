package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	m := sync.Mutex{}

	wg := &sync.WaitGroup{}

	count := func(wg *sync.WaitGroup) {
		for i := 0; i < 1000; i++ {
			m.Lock()
			counter++
			m.Unlock()
		}
		wg.Done()
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go count(wg)
	}
	wg.Wait()

	fmt.Println("Final value:", counter)
}
