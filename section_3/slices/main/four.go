package main

import "fmt"

func main() {
	src := []int{}
	src = append(src, 0)
	src = append(src, 1)
	src = append(src, 2)
	dest1 := append(src, 3)
	dest2 := append(src, 4)
	fmt.Println(src, dest1, dest2)
}
