package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type poolLocalInternal struct {
	pint       *int           //8
	p          unsafe.Pointer // 8
	private    interface{}    // 8(*itab) + 8(unsafe.Pointer) = 16
	shared     []interface{}  // 8(unsafe.Pointer) + 8(int) + 8(int) =24
	sync.Mutex                // 4(int32) + 4(uint32) = 8
}

func main() {
	fmt.Println(unsafe.Sizeof(poolLocalInternal{}))

}
