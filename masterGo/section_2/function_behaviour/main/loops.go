package main

import (
	"fmt"
	"unicode"
)

func main() {
	word := "abcde"
	for _, char := range word {
		upperWord := unicode.ToUpper(char)
		fmt.Print(string(upperWord))
	}
	fmt.Println("\n" + word)
}
