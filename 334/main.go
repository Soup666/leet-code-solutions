package main

import "fmt"

func main() {
	fmt.Println("Hello, 334!")

	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}

func increasingTriplet(nums []int) bool {
	m1 := 10000000
	m2 := m1

	for _, k := range nums {
		if k <= m1 {
			m1 = k
		} else if k <= m2 {
			m2 = k
		} else {
			return true
		}
	}

	return false
}

