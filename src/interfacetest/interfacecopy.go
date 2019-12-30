package main

import (
	"fmt"
)

type Rect struct {
	Width  int
	Height int
}

func main() {
	var a interface{}
	var r = Rect{50, 50}
	fmt.Printf("%p\n", &r)
	a = r
	fmt.Printf("%p\n", &a)
	var rx = a.(Rect)
	fmt.Printf("%p\n", &rx)
	r.Width = 100
	r.Height = 100
	fmt.Println(rx)
	//{50 50}

	//var a interface {}
	//var r = Rect{50, 50}
	//a = &r // 指向了结构体指针
	//
	//var rx = a.(*Rect) // 转换成指针类型
	//r.Width = 100
	//r.Height = 100
	//fmt.Println(rx)
	//&{100 100}
}
