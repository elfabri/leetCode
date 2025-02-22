/*
Difficulty: Hard
https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/description/

We run a preorder depth-first search (DFS) on the root of a binary tree.

At each node in this traversal, we output D dashes
(where D is the depth of this node), then we output the value of this node.
If the depth of a node is D, the depth of its immediate child is D + 1.
The depth of the root node is 0.

If a node has only one child, that child is guaranteed to be the left child.

Given the output traversal of this traversal, recover the tree and return its root.

it passed!
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
    arr := strings.Split(traversal, "-")
    idx := 0
    return traverse(&arr, 0, &idx)
}

func traverse(arr *[]string, depth int, idx *int) *TreeNode {
    if *idx >= len(*arr) {
        return nil
    }

    checkDepth := 0

    for (*arr)[*idx] == "" {
        checkDepth++
        *idx++
    }

    if checkDepth != depth - 1 && depth != 0 {
        *idx -= checkDepth
        return nil
    }

    var tree TreeNode

    tree.Val, _ = strconv.Atoi((*arr)[*idx])
    depth++
    *idx++

    tree.Left = traverse(arr, depth, idx)

    tree.Right = traverse(arr, depth, idx)

    return &tree
}

func main() {
    str := "1-2--3--4-5--6---7"
    fmt.Printf("trying with: %s\n", str)
    tree := recoverFromPreorder(str)
    /*
        1
      /   \
     2     5
    / \   /
   3   4 6
        /
       7
    */
    fmt.Printf(`
      %v
    /   \
   %v      %v
  / \    /
 %v   %v  %v
       /
      %v

`,

    tree.Val,
    tree.Left.Val,
    tree.Right.Val,
    tree.Left.Left.Val,
    tree.Left.Right.Val,
    tree.Right.Left.Val,
    tree.Right.Left.Left.Val,
    )
}
