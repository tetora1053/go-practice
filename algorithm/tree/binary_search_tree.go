// 二分探索木：挿入,検索
// 螺旋本p209,p214
package main

import (
	"fmt"
)

type node struct {
	key     int
	p, l, r *node
}

var (
	root *node
)

func insert(key int) {
	var (
		parent, cur, target *node
	)

	target = &node{key: key}
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

func find(key int) {
	cur := root

	for cur != nil {
		if cur.key == key {
			fmt.Println("yes")
			return
		} else if key < cur.key {
			cur = cur.l
		} else {
			cur = cur.r
		}
	}
	fmt.Println("no")
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
		switch cmd {
		case "print":
			inorder(root)
			fmt.Print("\n")
			preorder(root)
			fmt.Print("\n")
		case "insert":
			fmt.Scan(&key)
			insert(key)
		case "find":
			fmt.Scan(&key)
			find(key)
		}
	}
}
