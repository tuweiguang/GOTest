package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			ch <- 1
			fmt.Println("send ", i)
		}(i)
	}

	<-ch
	time.Sleep(time.Second * 5)
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second * 5)
	fmt.Println("exit!")
}
