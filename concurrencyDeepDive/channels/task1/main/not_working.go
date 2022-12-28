package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1024
	fmt.Println(<-c)
}
