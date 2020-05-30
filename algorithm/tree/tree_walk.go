// 螺旋本p198
// 木の巡回
// 先行順巡回(Preorder Tree Walk): id -> left -> right
// 中間順巡回(Inorder Tree Walk): left -> id -> right
// 後行順巡回(Postorder Tree Walk): left -> right -> id
package main

import (
	"fmt"
)

type node struct {
	p, l, r int
}

func preorder(tree []node, id int) {
	if id == -1 {
		return
	}
	fmt.Printf("%d ", id)
	preorder(tree, tree[id].l)
	preorder(tree, tree[id].r)
}

func inorder(tree []node, id int) {
	if id == -1 {
		return
	}
	if tree[id].l != -1 {
		inorder(tree, tree[id].l)
	}
	fmt.Printf("%d ", id)
	if tree[id].r != -1 {
		inorder(tree, tree[id].r)
	}
}

func postorder(tree []node, id int) {
	if id == -1 {
		return
	}
	if tree[id].l != -1 {
		postorder(tree, tree[id].l)
	}
	if tree[id].r != -1 {
		postorder(tree, tree[id].r)
	}
	fmt.Printf("%d ", id)
}

func main() {
	var n, id, l, r int
	fmt.Scan(&n)
	tree := make([]node, n)
	for i := 0; i < n; i++ {
		tree[i].p = -1
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&id, &l, &r)
		tree[id].l, tree[id].r = l, r
		if l != -1 {
			tree[l].p = id
		}
		if r != -1 {
			tree[r].p = id
		}
	}

	fmt.Println("Preorder")
	preorder(tree, 0)
	fmt.Print("\n")

	fmt.Println("Inorder")
	inorder(tree, 0)
	fmt.Print("\n")

	fmt.Println("Postorder")
	postorder(tree, 0)
	fmt.Print("\n")
}
