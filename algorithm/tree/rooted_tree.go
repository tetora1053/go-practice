// 螺旋本p188
// 根付き木
// ノード：親ノード、最も左の子ノード、右隣のノード
package main

import (
	"fmt"
)

type node struct {
	p, l, r int
}

func print(tree []node) {
	for i := 0; i < len(tree); i++ {
		fmt.Printf("node %d: parent = %d, depth = %d, ", i, tree[i].p, rec(tree, i, 0))
		if tree[i].p == -1 {
			fmt.Print("root, ")
		} else if tree[i].l == 0 {
			fmt.Print("leaf, ")
		} else {
			fmt.Print("internal node, ")
		}
		fmt.Print("[")
		for c := tree[i].l; c != 0; c = tree[c].r {
			fmt.Print(c)
			fmt.Print(",")
		}
		fmt.Print("]\n")
	}
}

func rec(tree []node, n, depth int) int {
	if tree[n].p == -1 {
		return depth
	}
	depth++
	depth = rec(tree, tree[n].p, depth)
	return depth
}

func main() {
	var n, id, k, c, l int
	fmt.Scan(&n)
	tree := make([]node, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&id, &k)
		if i == 0 {
			tree[i].p = -1
		}
		for j := 0; j < k; j++ {
			fmt.Scan(&c)
			if j == 0 {
				tree[i].l = c
				l = c
			} else {
				tree[l].r = c
			}
			tree[c].p = i
			l = c
		}
	}

	print(tree)
}
