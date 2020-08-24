package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 11
	b := *(*int)(unsafe.Pointer(&a)) // 新copy一个int变量b，并且b=11
	fmt.Println(&a, &b)
	fmt.Println(a, b)
	a = 111
	fmt.Println(a, b)

	c := 12
	d := (*int)(unsafe.Pointer(&c)) // 新的一个*int指针d，
	fmt.Println(&c, d)
	c = 122
	fmt.Println(c, *d)
}
