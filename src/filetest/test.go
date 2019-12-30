package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("failed to open file")
		os.Exit(2)
	}
	defer f.Close()

	for cnt := 0; cnt < 4000; cnt++ {
		content := fmt.Sprintf("stat,type=login name=%v %v\n", cnt, time.Now().UnixNano())
		if _, err := f.WriteString(content); err != nil {
			fmt.Printf("failed to write")
			break
		}
	}

}
