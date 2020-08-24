package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a int32 = 10

	atomic.StoreInt32(&a, 11)
	atomic.StoreInt32(&a, 12)
	fmt.Println(a)
}
