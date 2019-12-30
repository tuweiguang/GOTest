package main

import "fmt"

func main() {
	a := 10
	b := 10
	defer fmt.Println("defer1")
	defer func(a int) {
		//if err := recover(); err != nil {
		//	log.Printf("recover: %v", err)
		//}
		fmt.Println("defer2", a)
	}(a)
	defer func() {
		//if err := recover(); err != nil {
		//	log.Printf("recover: %v", err)
		//}
		fmt.Println("defer3", b)
	}()

	b = 20

	fmt.Println("main")

	//panic("EDDYCJY.")

	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Printf("recover: %v", err)
	//	}
	//}()
	//
	//panic("EDDYCJY.")
}
