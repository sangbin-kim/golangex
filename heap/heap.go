package main

import (
	"fmt"

	dataStruct "github.com/sangbin-kim/golangex/dataStruct"
)

func main() {

	h := &dataStruct.Heap{}

	// [-1, 3, -2, 5, 4] second largest

	nums := []int{-1, 3, -2, 5, 4}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 2 {
			h.Pop()
		}
	}
	fmt.Println(h.Pop())
}
