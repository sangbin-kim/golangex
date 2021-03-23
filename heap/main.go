package main

import (
	"fmt"

	dataStruct "github.com/sangbin-kim/golangex/dataStruct"
)

func main() {
	h := &dataStruct.Heap{}

	h.Push(7)
	h.Push(5)
	h.Push(6)
	h.Push(9)
	h.Push(8)
	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}
