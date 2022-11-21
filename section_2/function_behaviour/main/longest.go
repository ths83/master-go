package main

import "fmt"

func longest(strings ...string) int {
	var size int
	for _, s := range strings {
		sSize := len(s)
		if sSize > size {
			size = sSize
		}
	}

	return size
}

func main() {
	fmt.Println(longest("Six", "sleek", "swans", "swam", "swiftly", "southwards"))
}
