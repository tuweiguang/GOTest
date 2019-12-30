package main

import (
	"fmt"
	"unsafe"
)

type Base struct {
	name int
	age  int
}

func main() {
	//a := &Base{name:12,age:16}

	fmt.Println(unsafe.Sizeof(Base{}))
}
