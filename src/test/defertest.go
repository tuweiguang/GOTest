package main

import (
	"fmt"
)

func main() {
	fmt.Println(f())
	fmt.Println(f1())
	fmt.Println(f2())
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f1() (r int) {
	t := 5 //这里的t在defer func函数里面相当于全局变量
	defer func() {
		t = t + 5
	}()
	return t

}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r) //r作为值传递，形参会重新创建变量r
	return 1
}

/* 运行结果
1
5
1
*/

/*
结论：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
*/
