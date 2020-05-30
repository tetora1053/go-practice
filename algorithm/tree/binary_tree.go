// 螺旋本p193
// 二分木
package main

import (
	"fmt"
)

type node struct {
	p, l, r int
}

func print(tree []node) {
	for i := 0; i < len(tree); i++ {
		fmt.Printf("node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s\n", i, tree[i].p, getSibling(tree, i), getDegree(tree, i), getDepth(tree, i), getHeight(tree, i), getType(tree, i))
	}
}

func getSibling(tree []node, id int) int {
	if tree[id].p == -1 {
		return -1
	}
	if tree[tree[id].p].l == id {
		return tree[tree[id].p].r
	} else {
		return tree[tree[id].p].l
	}
}

func getDegree(tree []node, id int) (degree int) {
	if tree[id].l != -1 {
		degree++
	}
	if tree[id].r != -1 {
		degree++
	}
	return
}

func getDepth(tree []node, id int) int {
	if tree[id].p == -1 {
		return 0
	}
	return 1 + getDepth(tree, tree[id].p)
}

func getHeight(tree []node, id int) int {
	if tree[id].l == -1 && tree[id].r == -1 {
		return 0
	}
	if 1+getHeight(tree, tree[id].l) >= 1+getHeight(tree, tree[id].r) {
		return 1 + getHeight(tree, tree[id].l)
	} else {
		return 1 + getHeight(tree, tree[id].r)
	}
}

func getType(tree []node, id int) string {
	if tree[id].p == -1 {
		return "root"
	} else if tree[id].l == -1 && tree[id].r == -1 {
		return "leaf"
	}
	return "internal node"
}

func main() {
	var n, v, l, r int
	fmt.Scan(&n)
	tree := make([]node, n)
	for i := 0; i < n; i++ {
		tree[i].p = -1
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&v, &l, &r)
		tree[v].l, tree[v].r = l, r
		if l != -1 {
			tree[l].p = v
		}
		if r != -1 {
			tree[r].p = v
		}
	}
	print(tree)
}
