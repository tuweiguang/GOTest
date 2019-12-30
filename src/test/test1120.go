package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.WaitGroup

	for i := 0; i < 5; i++ {
		scene.Add(1)
		go func(w sync.WaitGroup, i int) {
			fmt.Println(i)
			w.Done()
		}(scene, i)
	}

	scene.Wait()
	fmt.Println("exit")

}
