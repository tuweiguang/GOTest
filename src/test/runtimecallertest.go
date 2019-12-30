package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func f12(depth string) {
	d, _ := strconv.Atoi(depth)
	rpc := make([]uintptr, 1)
	n := runtime.Callers(d+1, rpc)
	if n < 1 {
		return
	}
	frame, _ := runtime.CallersFrames(rpc).Next()
	fmt.Printf("%v %v %v\n", frame.File, frame.Function, frame.Line)
}

func f11(depth string) (f string, l int) { //calldepth = 0
	d, _ := strconv.Atoi(depth)
	_, file, line, ok := runtime.Caller(d)
	if !ok {
		file = "???"
		line = 0
	}
	return file, line
}

func f22(depth string) (f string, l int) { //calldepth = 1
	return f11(depth)
}

func main() { //calldepth = 2
	f, l := f22(os.Args[1])
	fmt.Printf("%v %v\n", f, l)

	f12(os.Args[1])
}

/*
go run test.go 0
XXX/test.go 12

go run test.go 1
XXX/test.go 21

go run test.go 2
XXX/test.go 25

go run test.go 3
E:/Go/src/runtime/proc.go 198
*/
