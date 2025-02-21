package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 238!")

	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	n := len(nums)

	left[0] = 1

	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := make([]int, len(nums))

	right[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		left[i] = left[i] * right[i]
	}

	return left
}

