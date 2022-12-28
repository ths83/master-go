package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan int)

	go write(c, done)
	go read(1, c)
	go read(2, c)

	<-done
}

func write(c chan<- int, done chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}

	close(done)
}

func read(n int, c <-chan int) {
	for {
		fmt.Printf("Reader %d reads %d\n", n, <-c)
	}
}
