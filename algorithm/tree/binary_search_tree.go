// 二分探索木：挿入,検索、削除
// 螺旋本p209,p214,p217
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

func delete(key int) {
	var target, c *node
	cur := root

	for cur != nil {
		if cur.key == key {
			// 削除対象ノードがリーフノード、もしくは子を一つしか持たない場合
			if cur.l == nil || cur.r == nil {
				target = cur
				// 削除対象ノードが子を二つ持っている場合
			} else {
				target = getSuccessor(cur.r)
				cur.key = target.key
			}

			if target.l != nil {
				c = target.l
			} else if target.r != nil {
				c = target.r
			}
			break
		}
		if key < cur.key {
			cur = cur.l
		} else {
			cur = cur.r
		}
	}

	// 削除対象ノードがなければreturn
	if target == nil {
		return
	}

	if c != nil {
		c.p = target.p
	}
	if target == target.p.l {
		target.p.l = c
	} else {
		target.p.r = c
	}
}

func getSuccessor(node *node) *node {
	cur := node
	for cur != nil {
		cur = cur.l
	}
	return cur
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
		case "delete":
			fmt.Scan(&key)
			delete(key)
		}
	}
}
