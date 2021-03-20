package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	list := &dataStruct.LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNode()
	list.RemoveNode(list.Root.Next)
	list.PrintNode()

	list.RemoveNode(list.Root)
	list.PrintNode()

	list.RemoveNode(list.Tail)
	list.PrintNode()

	list.PrintReverse()

	a := []int{1, 2, 3, 4, 5}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)
}
