package main

import (
	"fmt"
	"strings"
)

func count(s string, m map[string]int) {
	words := strings.Trim(s, " ")
	for _, word := range words {
		m[string(word)]++
	}
}

func printCount(m map[string]int) {
	for key := range m {
		if m[key] > 1 {
			fmt.Println(key)
		}
	}
}
