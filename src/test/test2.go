package main

import "fmt"

func main() {
	a := 10
	b := 10

	defer func() {

	}()
	defer func(a int) {
		fmt.Println("defer2", a)
	}(a)
	defer func() {
		fmt.Println("defer3", b)
	}()

	b = 20
	//fmt.Println("main")
}
