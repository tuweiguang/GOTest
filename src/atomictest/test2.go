package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var read atomic.Value

	m1 := map[string]interface{}{
		"a": 1,
		"b": 2,
	}

	read.Store(m1)
	fmt.Println(read)

	m2 := map[string]interface{}{
		"c": 3,
		"d": 4,
	}
	read.Store(m2)
	fmt.Println(read)

}
