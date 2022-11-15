package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func acronym(s string) (acr string) {

	var acronym string

	for _, r := range s {
		if unicode.IsUpper(r) {
			acronym += string(r)
		}
	}

	return acronym
}

func main() {
	s := "Pan Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronym(s))
}
