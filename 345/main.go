package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	fmt.Println("Hello, 345!")

	args := os.Args[1:]

	if (len(args) < 1) {
		fmt.Println("Usage: ./main.go <input>")
		return
	}

	fmt.Println(reverseVowels(args[0]))
}


var VOWELS = []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func reverseVowels(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !slices.Contains(VOWELS, runes[left]) {
			left++
		}
		for left < right && !slices.Contains(VOWELS, runes[right]) {
			right--
		}
		runes[left], runes[right] = runes[right], runes[left]

		left++
		right--
	}

	return string(runes)
}
