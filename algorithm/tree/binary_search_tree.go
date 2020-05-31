// 二分探索木：挿入
// 螺旋本p209
package main

import (
	"fmt"
)

type node struct {
	key     int
	p, l, r *node
}

var (
	cur  *node
	root *node
)

func insert(key int) {
	var parent *node

	target := &node{key: key}
	cur = root

	for cur != nil {
		parent = cur
		if target.key < cur.key {
			cur = cur.l
		} else {
			cur = cur.r
		}
	}
	target.p = parent

	if parent == nil {
		root = target
	} else {
		if target.key < parent.key {
			parent.l = target
		} else {
			parent.r = target
		}
	}
}

func inorder(root *node) {
	if root == nil {
		return
	}
	inorder(root.l)
	fmt.Printf("%d ", root.key)
	inorder(root.r)
}

func preorder(root *node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.key)
	preorder(root.l)
	preorder(root.r)
}

func main() {
	var (
		cmd    string
		n, key int
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cmd)
		if cmd == "print" {
			inorder(root)
			fmt.Print("\n")
			preorder(root)
			fmt.Print("\n")
		} else {
			fmt.Scan(&key)
			insert(key)
		}
	}
}
