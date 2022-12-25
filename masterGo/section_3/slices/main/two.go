package main

import "fmt"

func changeSlice1(s []int) {
	s[0] = 7
}

func main() {
	s1 := []int{1}
	fmt.Println("s1 before changeSlice1:", s1)
	changeSlice1(s1)
	fmt.Println("s1 after changeSlice1:", s1)
}
