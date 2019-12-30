package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintPanicInfo() []string {
	var name, file string
	var line int
	var pc [16]uintptr

	result := make([]string, 0)

	n := runtime.Callers(3, pc[:])
	fmt.Println("=======", n)
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		result = append(result, fmt.Sprintf("--- file:%v,  method:%v   line:%v --", file, name, line))
	}

	return result
}

func RPCPanicCatch() {
	x := recover()
	if x != nil {
		ee := PrintPanicInfo()
		fmt.Printf("Panic Catch :%v", x)
		for i, v := range ee {
			fmt.Printf("CallStack %v:%v", i, v)
		}
		//RPCCleanAndExit()
	}
}

func test() {
	defer RPCPanicCatch()
	panic("test")
}
func main() {
	go test()
	for {
		time.Sleep(time.Second)
	}
}
