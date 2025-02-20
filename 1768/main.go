
package main

import (
  "fmt"
  "os"
)

func main() {
  args := os.Args[1:]
  
  if (len(args) < 2) {
    fmt.Println("Usage: ./main.go <word1> <word2>")
    return
  }

  fmt.Println(mergeAlternately(args[0], args[1]))
}

func mergeAlternately(word1 string, word2 string) string {
  
  smallWordLen := min(len(word1), min(len(word2)))
  result := ""

  words := []string{word1, word2}

  for i := 0; i < smallWordLen*2; i += 1 {
    result += string(words[(i % 2)][(i / 2)])
  }

  return result + word1[smallWordLen:] + word2[smallWordLen:]
}
