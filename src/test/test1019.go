package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 1000
	s := "hello world"

	//任何类型的指针可以转成unsafe.Pointer。
	pa1 := unsafe.Pointer(&a)
	ps1 := unsafe.Pointer(&s)
	fmt.Println(&a)  //0xc042062080
	fmt.Println(&s)  //0xc0420561c0
	fmt.Println(pa1) //0xc042062080
	fmt.Println(ps1) //0xc0420561c0

	//unsafe.Pointer可以转成任何类型的指针。
	pa2 := (*int32)(unsafe.Pointer(&a)) //int32占4字节
	fmt.Println(*pa2)                   //1000 1111101000
	pa3 := (*byte)(unsafe.Pointer(&a))  //byte占1字节
	fmt.Println(*pa3)                   //232    11101000  截取第1字节的数据

	//指针地址不能直接进行数学运算，要转成uintptr类型进行数学运算
	arr := [3]int{1, 2, 3}
	ps2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + unsafe.Sizeof(int(0))))
	fmt.Println(*ps2) //2
}
