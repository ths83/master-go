package main

import (
	"fmt"
)

func worker(done chan int) {
	// Do something for some time
	for i := 0; i < 1000; i++ {
		fmt.Print(".")
	}
	fmt.Println()
	close(done)
}

func main() {
	done := make(chan int)
	// start the goroutine.
	go worker(done)

	fmt.Println("Waiting for the goroutine")
	<-done
	fmt.Println("Done")
}
