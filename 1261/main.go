package main

import (
  "fmt"
  "os"
)

type FindElements struct {
    m map[int]bool
}


func Constructor(root *TreeNode) FindElements {
    m := make(map[int]bool)
    solve(root, 0, m)
    return FindElements{
        m: m,
    }
}

func solve(root *TreeNode, val int, m map[int]bool) {
    if root == nil {
        return
    }

    m[val] = true
    solve(root.Left, 2*val+1, m)
    solve(root.Right, 2*val+2, m)
}


func (this *FindElements) Find(target int) bool {
    return this.m[target]
}


