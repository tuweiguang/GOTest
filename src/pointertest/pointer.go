package main

import (
	"fmt"
)

type A struct {
	aa int
}

func main() {
	//new 分配出来的是指针
	a := new(int)
	var b int = 1
	a = &b
	fmt.Printf("%v\n", *a)

	//make
	c := make([]int, 10)
	var d []int
	c = d
	fmt.Println(c)

	//
	e := A{aa: 1}
	f := &A{aa: 2}
	var g A
	var h *A

	var j *A
	g = e
	h = f
	j = &e

	fmt.Println(g, h, j)

}
