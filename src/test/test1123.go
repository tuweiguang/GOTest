package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type A struct {
	methodLocker sync.Mutex
	v            int
}

func main() {
	m := make(map[string]*A)
	a := &A{v: 100}
	m["a"] = a

	go func() {
		for {
			time.Sleep(time.Second * 5)
		}
	}()

	fmt.Println("-------0------")
	l := m["a"]
	l.methodLocker.Lock()
	fmt.Println("-------1------")

	go func() {
		time.Sleep(time.Second * 5)
		delete(m, "a")
		l = nil
		runtime.GC()
	}()

	fmt.Println("-------2------")
	//l = m["a"]
	l.methodLocker.Lock()
	fmt.Println("-------3------")

}
