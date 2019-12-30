package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	select { //select 没有阻塞和default时候，只会执行一个case语句
	case <-ch1:
		fmt.Printf("ch1\n")
	case <-ch2:
		fmt.Printf("ch2\n")
	}
}

/*
运行：
ch1
或者
ch2
*/
