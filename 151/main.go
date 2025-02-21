package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Hello, 151!")

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: ./main.go <string>")
		return
	}

	fmt.Println(reverseWords(args[0]))
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	slices.Reverse(words)
	return strings.Join(removeWhiteSpace(words), " ")
}

func removeWhiteSpace(s []string) []string {
	trimmedS := make([]string, 0)
	for _, word := range s {
		if word != "" {
			trimmedS = append(trimmedS, word)
		}
	}
	return trimmedS
}
