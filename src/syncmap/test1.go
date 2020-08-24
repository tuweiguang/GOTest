package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map //不需要初始化

	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", "xxx")

	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	fmt.Println(scene.Load("egypt"))

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
