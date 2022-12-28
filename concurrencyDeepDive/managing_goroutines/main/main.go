package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker(i, &wg)
	}
}

func worker(n int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < rand.Intn(10)+10; i++ {
		fmt.Print(n)
		time.Sleep(100 * time.Millisecond)
	}
}
