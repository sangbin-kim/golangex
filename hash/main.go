package main

import (
	"fmt"

	dataStruct "github.com/sangbin-kim/golangex/dataStruct"
)

func main() {
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcdf = ", dataStruct.Hash("abcdf"))
	fmt.Println("tbcde = ", dataStruct.Hash("tbcde"))
	fmt.Println("abcdedfsfs = ", dataStruct.Hash("abcdedfsfs"))

	m := dataStruct.CreateMap()
	m.Add("AAA", "01090042504")
	m.Add("BBB", "01090041504")
	m.Add("CDJDLSLFIEJH", "01090023304")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("BBB = ", m.Get("BBB"))
	fmt.Println("CDJDLSLFIEJH = ", m.Get("CDJDLSLFIEJH"))
}
