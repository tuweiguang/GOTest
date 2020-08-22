package main

import "fmt"

type T1 struct {
	a int
	b int
}

type T2 struct {
	a string
	b string
}

type T3 struct {
	a int
	b int
}

func main() {
	v1 := T1{a: 1, b: 1}
	v2 := T2{a: "111", b: "111"}
	v3 := T3{a: 2, b: 2}
	v4 := T1{a: 2, b: 2}

	fmt.Println(v1 == v2)

	fmt.Println(v1 == T1(v3))

	fmt.Println(v1 == v4)
}
