package dataStruct

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.Root == nil {
		l.Root = &Node{Val: val}
		l.Tail = l.Root
		return
	}
	l.Tail.Next = &Node{Val: val} //node 생성하여 현 tail의 next 포인터로 지정
	prev := l.Tail                // 기존 tail을 prev로 지정
	l.Tail = l.Tail.Next          // 새로운 tail을 새로만들어진 node 포인터로 지정
	l.Tail.Prev = prev            // 새로운 tail의 prev를 기존 tail로 지정

}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next
		if l.Root != nil {
			l.Root.Prev = nil
		}

		node.Next = nil
		return
	}
	prev := node.Prev
	// 지우고자 하는 node가 tail인 경우
	if node == l.Tail {
		prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev

	} else { // 지우고자 하는 node가 root 또는 tail이 아닌 경우
		node.Prev = nil
		prev.Next = prev.Next.Next
		prev.Next.Prev = prev
	}
	node.Next = nil
}

func (l *LinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}

func (l *LinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func (l *LinkedList) Empty() bool {
	return l.Root == nil
}

func (l *LinkedList) PrintNode() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

func (l *LinkedList) Front() int {
	if l.Root != nil {
		return l.Root.Val
	}
	return 0
}

func (l *LinkedList) PopFront() {
	if l.Root == nil {
		return
	}
	l.RemoveNode(l.Root)
}

func (l *LinkedList) PrintReverse() {
	node := l.Tail
	for node.Prev != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}
	fmt.Printf("%d\n", node.Val)
}
