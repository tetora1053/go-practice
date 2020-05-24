package main

import (
	"fmt"
)

type node struct {
	prev, next *node
	key        int
}

var (
	nilNode node
)

func init() {
	nilNode.prev = &nilNode
	nilNode.next = &nilNode
}

func insert(node node) {
	nilNode.next.prev = &node
	node.next = nilNode.next
	node.prev = &nilNode
	nilNode.next = &node
}

func delete(node *node) {
	if node == &nilNode {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func deleteFirst() {
	delete(nilNode.next)
}

func deleteLast() {
	delete(nilNode.prev)
}

func searchNode(key int) *node {
	cur := nilNode.next
	for cur != &nilNode && cur.key != key {
		cur = cur.next
	}
	return cur
}

func main() {
	var (
		n, key int
		cmd    string
	)

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cmd)
		switch cmd {
		case "insert":
			fmt.Scan(&key)
			node := node{key: key}
			insert(node)
		case "delete":
			fmt.Scan(&key)
			delete(searchNode(key))
		case "deleteFirst":
			deleteFirst()
		case "deleteLast":
			deleteLast()
		}
	}

	cur := nilNode.next
	for {
		if cur != &nilNode {
			fmt.Println(cur.key)
			cur = cur.next
		}
	}
}
