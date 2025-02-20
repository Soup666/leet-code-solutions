package main

import (
  "fmt"
  "os"
)

func main() {
  args := os.Args[1:]

  if (len(args) < 2) {
    fmt.Println("Usage: ./main2.go <word1> <word2>")
    return
  }

  fmt.Println(gcdOfStrings(args[0], args[1]))
}

func gcdOfInts(a int, b int) int {

  for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
  gcd := gcdOfInts(len(str1), len(str2))

  if (str1 + str2 == str2 + str1) {
    return str1[:gcd]
  } else {
    return ""
  }
}
