package main

import (
	"os"
	"strconv"
)

func fizzbuzz(n int) {
	var fizz = "fizz"
	var buzz = "buzz"
	var fizzbuzz = fizz + buzz

	switch {
	case n%3 == 0:
		print(fizz)
	case n%5 == 0:
		print(buzz)
	case n%15 == 0:
		print(fizzbuzz)
	default:
		print(n)
	}
}
func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
