/*
2467. Most Profitable Path in a Tree
Medium

There is an undirected tree with n nodes
labeled from 0 to n - 1, rooted at node 0.

You are given a 2D integer array edges of length
n - 1 where edges[i] = [ai, bi] indicates that
there is an edge between nodes ai and bi in the tree.

At every node i, there is a gate. You are also
given an array of even integers amount,
where amount[i] represents:

    * the price needed to open the gate
      at node i, if amount[i] is negative, or,

    * the cash reward obtained on opening the gate at node i, otherwise.

The game goes on as follows:

    * Initially, Alice is at node 0 and Bob is at node bob.

    * At every second, Alice and Bob each move
      to an adjacent node. Alice moves towards some leaf node,
      while Bob moves towards node 0.

    * For every node along their path, Alice and Bob
      either spend money to open the gate
      at that node, or accept the reward. Note that:

        * If the gate is already open, no price
          will be required, nor will there be any cash reward.

        * If Alice and Bob reach the node simultaneously,
          they share the price/reward for opening
          the gate there. In other words, if the price
          to open the gate is c, then both
          Alice and Bob pay c / 2 each. Similarly,
          if the reward at the gate is c,
          both of them receive c / 2 each.

    * If Alice reaches a leaf node, she stops moving.
      Similarly, if Bob reaches node 0, he stops moving.
      Note that these events are independent of each other.

Return the maximum net income Alice can have if she travels towards the optimal leaf node.

Constraints:

    2 <= n <= 105

    edges.length == n - 1

    edges[i].length == 2

    0 <= ai, bi < n

    ai != bi

    edges represents a valid tree.

    1 <= bob < n

    amount.length == n

    amount[i] is an even integer in the range [-104, 104].

*/

package main

import "fmt"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {

}

func main() {
    fmt.Printf("starting with edges = [[0,1],[1,2],[1,3],[3,4]], bob = 3, amount = [-2,4,2,-4,6]\n")
    edges := [][]int{{0,1},{1,2},{1,3},{3,4}}
    bob := 3
    amount := []int{-2,4,2,-4,6}
    result := mostProfitablePath(edges, bob, amount)
    fmt.Printf("Should be 6, got: %d\n", result)
}
