package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, 1028!")

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: go run main.go <input>")
		return
	}

	printPreorder(recoverFromPreorder(args[0]))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printPreorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printPreorder(root.Left)
	printPreorder(root.Right)
}

func recoverFromPreorder(traversal string) *TreeNode {

	var stack []*TreeNode
	n := len(traversal) // length of input
	i := 0              // current index parsing

	for i < n {
		// Count the number of leading dashes (depth)
		depth := 0
		for i < n && traversal[i] == '-' {
			depth++
			i++
		}

		// Extract the node value
		start := i
		for i < n && traversal[i] >= '0' && traversal[i] <= '9' {
			i++
		}
		value, _ := strconv.Atoi(traversal[start:i])

		// Create new node
		node := &TreeNode{Val: value}

		// Attach node to its parent in the stack
		if len(stack) > 0 {
			if depth > len(stack)-1 {
				stack[len(stack)-1].Left = node
			} else {
				stack = stack[:depth]
				stack[len(stack)-1].Right = node
			}
		}

		// Push current node onto stack
		stack = append(stack, node)
	}

	// Root is at index 0 in the stack
	return stack[0]
}

