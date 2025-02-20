package main

import (
  "fmt"
  "os"
  "strconv"
  "slices"
  "strings"
)

func main() {
  args := os.Args[1:]

  if (len(args) < 2) {
    fmt.Println("Usage ./main.go \"<int array>\" <int>")
    return
  }
  
  nums, err := parseIntList(args[0])

  if (err != nil) {
    fmt.Println("Error: ", err)
    return
  }

  num, err := strconv.Atoi(args[1])

  if (err != nil) {
    fmt.Println("Error: ", err)
    return
  }

  fmt.Println(kidsWithCandies(nums, num))

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

func kidsWithCandies(candies []int, extraCandies int) []bool {
  maxCandie := slices.Max(candies)
  result := make([]bool, len(candies))

  for i, candy := range candies {
    result[i] = candy + extraCandies >= maxCandie
  }

  return result
}
