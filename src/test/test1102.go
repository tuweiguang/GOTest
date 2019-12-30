package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	//conn, err := jsonrpc.Dial("tcp", "182.61.200.7:80")
	conn, err := net.Dial("tcp", "182.61.200.7:80")
	if err != nil {
		fmt.Println("dial error:", err)
	}
	tmp := make([]byte, 256) // using small tmo buffer for demonstrating
	fmt.Println(time.Now(), "start read")
	for {
		_, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			fmt.Println(time.Now(), err)
			break
		}
		//buf = append(buf, tmp[:n]...)
		fmt.Println(time.Now(), tmp)
	}
}
