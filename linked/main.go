package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node
	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNode(root)
	root, tail = RemoveNode(root.next.next, root, tail)
	PrintNode(root)
}

func AddNode(tail *Node, val int) *Node {

	node := &Node{val: val}
	tail.next = node

	return node
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	// 지우고자 하는 node가 root인 경우
	if node == root {
		root = root.next
		return root, tail
	}
	// 아닌 경우 지우고자 하는 node 이전 node 찾기
	prev := root
	for prev.next != node {
		prev = prev.next
	}
	// 지우고자 하는 node가 tail인 경우
	if node == tail {
		prev.next = nil
		tail = prev

	} else { // 지우고자 하는 node가 root 또는 tail이 아닌 경우
		prev.next = prev.next.next
	}

	return root, tail
}

func PrintNode(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
