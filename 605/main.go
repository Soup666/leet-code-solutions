package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  args := os.Args[1:]

  if (len(args) < 2) {
    fmt.Println("Usage: ./main.go \"<flowerbed>\" n")
    return
  }

  flowerbed, err := parseIntList(args[0])

  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  n, err := strconv.Atoi(args[1])

  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  fmt.Println(canPlaceFlowers(flowerbed, n))
}

// Courtesy of ChatGPT
func parseIntList(s string) ([]int, error) {
        parts := strings.Split(s, ",") // Split by comma
        nums := make([]int, len(parts))

        for i, part := range parts {
                n, err := strconv.Atoi(strings.TrimSpace(part)) // Convert to int
                if err != nil {
                        return nil, err // Return error if parsing fails
                }
                nums[i] = n
        }

        return nums, nil
}

func canPlaceFlowers(flowerbed []int, n int) bool {

  if (n == 0) {
    return true
  }

  for i := 0; i < len(flowerbed); i++ {
    if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
      flowerbed[i] = 1
      n--

      if (n == 0) {
        return true
      }
    }
  }
  return false
}
