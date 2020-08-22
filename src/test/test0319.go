package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var v atomic.Value
	v.Store("joker")
	fmt.Println(v.Load()) // joker
	v.Store("bob")
	fmt.Println(v.Load())
}
