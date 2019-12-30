//跨语言RPC
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"sync"
	"time"
)

// 算数运算结构体
type Arith struct {
}

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

var m sync.Mutex
var count int

func (this *Arith) Mylock(req int, res *int) error {
	log.Println("mylock")

	m.Lock()

	return nil
}

func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	log.Println("Multiply")
	if count == 0 {
		time.Sleep(time.Second * 20)
		count++
	}
	res.Pro = req.A * req.B
	return nil
}

func (this *Arith) Myunlock(req int, res *int) error {
	log.Println("myunlock")
	m.Unlock()
	return nil
}

func main() {
	rpc.Register(new(Arith)) // 注册rpc服务
	lis, err := net.Listen("tcp4", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			line, _, err := reader.ReadLine()
			if err != nil {
				fmt.Printf("read command error:%v\n", err)
				continue
			}
			if string(line) == "q" {
				os.Exit(0)
			}
		}
	}()

	fmt.Fprintf(os.Stdout, "%s\n", "start connection")
	for {
		conn, err := lis.Accept() // 接收客户端连接请求
		if err != nil {
			continue
		}

		go func(conn net.Conn) { // 并发处理客户端请求
			fmt.Fprintf(os.Stdout, "%s,my add:%v peer add %v\n", "new client in coming", conn.LocalAddr(), conn.RemoteAddr())
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
