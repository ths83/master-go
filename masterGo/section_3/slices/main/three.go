package main

import "fmt"

func changeSlice2(s []int) {
	s = []int{7}
}

func main() {
	s1 := []int{1}
	fmt.Println("s1 before changeSlice2:", s1)
	changeSlice2(s1)
	fmt.Println("s1 after changeSlice2:", s1)
}
