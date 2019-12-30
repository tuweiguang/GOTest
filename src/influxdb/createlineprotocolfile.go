package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("backup.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("failed to open file")
		os.Exit(2)
	}
	defer f.Close()

	t := time.Now().UnixNano()
	for cnt := 0; cnt < 4000; cnt++ {
		content := fmt.Sprintf("stat,type=login name=%v %v\n", cnt, t)
		if _, err := f.WriteString(content); err != nil {
			fmt.Printf("failed to write")
			break
		}
		t++
	}

}
