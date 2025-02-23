/*
889. Construct Binary Tree from Preorder and Postorder Traversal
Difficulty: medium

Given two integer arrays, preorder and postorder
where preorder is the preorder traversal of a binary tree
of distinct values and postorder is the postorder traversal
of the same tree, reconstruct and return the binary tree.

If there exist multiple answers, you can return any of them.

Constraints:

* 1 <= preorder.length <= 30
* 1 <= preorder[i] <= preorder.length
* All the values of preorder are unique.
* postorder.length == preorder.length
* 1 <= postorder[i] <= postorder.length
* All the values of postorder are unique.

It is guaranteed that preorder and postorder
are the preorder traversal and postorder traversal
of the same binary tree.
*/

package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    
    var tree TreeNode
    if len(preorder) == 1 {
        tree.Val = preorder[0]
        return &tree
    }

    last := len(preorder) - 1

    return constructTree(0, last, 0, &preorder, &postorder)
}

func constructTree(preStart, preEnd, postStart int, pre, post *[]int) *TreeNode {
    if preStart > preEnd {
        return nil
    }

    var tree TreeNode
    if preStart == preEnd {
        // tree contains only one node
        tree.Val = (*pre)[preStart]
        return &tree
    }

    // number of Nodes to be processed in the left arm
    leftRoot := (*pre)[preStart + 1]
    nodesInLeft := 1
    for leftRoot != (*post)[postStart + nodesInLeft - 1] {
        nodesInLeft++
    }

    // root
    tree.Val = (*pre)[preStart]

    tree.Left = constructTree(preStart + 1, preStart + nodesInLeft, postStart, pre, post)
    tree.Right = constructTree(preStart + nodesInLeft + 1, preEnd, postStart + nodesInLeft, pre, post)

    return &tree
}

func main() {
    fmt.Printf("pre: [1,2,4,5,3,6,7]\n")
    fmt.Printf("post: [4,5,2,6,7,3,1]\n")
    pre := []int{1,2,4,5,3,6,7}
    post := []int{4,5,2,6,7,3,1}

    tree := constructFromPrePost(pre, post)

    /*
    given solution: 
        1
      /   \
     2     3
    / \   / \
   4   5 6   7

    */

    fmt.Println("position_depth:")
    fmt.Printf("head_0: %v\n", tree.Val)
    fmt.Printf("Left_1: %v\n", tree.Left.Val)
    fmt.Printf("Right_1: %v\n", tree.Right.Val)
    fmt.Printf("Left-Left_2: %v\n", tree.Left.Left.Val)
    fmt.Printf("Left-Right_2: %v\n", tree.Left.Right.Val)
    fmt.Printf("Right-Left_2: %v\n", tree.Right.Left.Val)
    fmt.Printf("Right-Right_2: %v\n", tree.Right.Right.Val)
}
