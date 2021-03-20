package main

import "fmt"

func main() {
	a := make([]int, 0, 8)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	a = append(a, 1)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

}
