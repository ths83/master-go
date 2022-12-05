package main

import "fmt"

func appendOne(s *[]int) {
	*s = append(*s, 1)
}

func main() {
	//s1 := []int{0, 0, 0, 0}
	s1 := make([]int, 4, 8) // capacity is twice the initial size
	s2 := s1
	fmt.Printf("Before appendOne:\ns1: %v\ns2: %v\n", s1, s2)
	appendOne(&s1)
	fmt.Printf("After appendOne:\ns1: %v\ns2: %v\n", s1, s2)
	s1[0] = 2
	fmt.Printf("After changing s1:\ns1: %v\ns2: %v\n", s1, s2)

	appendOne(&s2)
	fmt.Printf("After changing s2:\ns1: %v\ns2: %v\n", s1, s2)
}
