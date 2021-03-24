package main

import "fmt"

type Key struct {
	v int
}

type Value struct {
	v int
}

func main() {
	var m map[string]string

	m = make(map[string]string)

	m["abc"] = "bbb"
	fmt.Println(m["abc"])

	m1 := make(map[int]string)
	m1[53] = "ddd"

	fmt.Println(m1[53])

	fmt.Println(m1[55])

	m2 := make(map[int]int)

	m2[4] = 4
	fmt.Println("m2[10] = ", m2[10])

	m3 := make(map[int]bool)
	m3[4] = true
	fmt.Println(m3[6])

	v, ok := m2[10]
	v1, ok1 := m2[4]

	fmt.Println(v, ok, v1, ok1)

	delete(m2, 4)

	m2[2] = 5
	m2[10] = 101

	for key, value := range m2 {
		fmt.Println(key, " = ", value)
	}
}
