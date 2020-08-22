package main

import (
	"fmt"
	"path"
)

func main() {
	r := path.Join("aa", "bb", "cc")
	fmt.Println(r)
}
