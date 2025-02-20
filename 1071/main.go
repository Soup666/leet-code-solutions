package main

import (
  "os"
  "fmt"
  "strings"
  "sort"
)

func main() {
  args := os.Args[1:]

  if (len(args) < 2) {
    fmt.Println("Usage: ./main.go <word1> <word2>")
    return
  }

  fmt.Println(gcdOfStrings(args[0], args[1]))
}

func gcdOfStrings(str1 string, str2 string) string {

  if (str1 == str2) {
    return str1
  }

  result := ""

  words := []string{str1, str2}

  sort.Slice(words, func(i, j int) bool {
    return len(words[i]) < len(words[j])
  })
  
  for i := 1; i <= len(words[1]) / 2; i++ {

    if (i > len(words[0])) {
      continue;
    }

    term := words[0][:i]

    if (strings.Repeat(term, (len(words[1]) / len(term))) == words[1] && strings.Repeat(term, (len(words[0]) / len(term))) == words[0]) {
      result = term
    }
  }

  return result
}
