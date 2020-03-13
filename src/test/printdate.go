package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("params error!")
	}

	start, ok := strconv.Atoi(os.Args[1])
	if ok != nil {
		fmt.Println("params error!")
	}
	end, ok := strconv.Atoi(os.Args[2])
	if ok != nil {
		fmt.Println("params error!")
	}

	f, ok := os.OpenFile("birthday.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if ok != nil {
		fmt.Println("openfile error!")
	}

	for i := start; i <= end; i++ {
		for j := 1; j <= 12; j++ {
			for k := 1; k <= 31; k++ {
				date := fmt.Sprintf("%d%02d%02d\n", i, j, k)
				f.Write([]byte(date))
			}
		}
		fmt.Println("%d finish!", i)
	}
	f.Close()
}
