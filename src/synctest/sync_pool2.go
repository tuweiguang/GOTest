package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//p := &sync.Pool{
	//	New: func() interface{} {
	//		return 0
	//	},
	//}
	//
	//a := p.Get().(int)
	//p.Put(1)
	//for i:=0;i< 100000;i++{
	//	obj := make([]byte, 1024)
	//	_ = obj
	//}
	//runtime.GC()
	//b := p.Get().(int)
	//fmt.Println(a, b)

	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	runtime.GOMAXPROCS(2)

	a := p.Get().(int)
	fmt.Println(a)
	p.Put(1)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		p.Put(100)
	}()
	wg.Wait()

	time.Sleep(time.Second * 1)

	p.Put(4)
	p.Put(5)

	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
}
