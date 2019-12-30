//跨语言RPC
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

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

func main() {
	var conn *rpc.Client
	var err error

	timeout := time.Duration(10 * time.Second)
	ch := make(chan struct{})
	go func() {
		conn, err = jsonrpc.Dial("tcp", "127.0.0.1:8096")
		if err != nil {
			log.Fatalln("dailing error: ", err)
		}

		req := ArithRequest{9, 2}
		var res ArithResponse
		fmt.Println(time.Now(), "first call")
		err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
		if err != nil {
			log.Println("arith error: ", err)
		}
		log.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
	}()

	go func() {
		<-ch
		time.Sleep(time.Second * 3)
		conn2, err2 := jsonrpc.Dial("tcp", "127.0.0.1:8096")
		if err2 != nil {
			log.Fatalln("dailing error: ", err2)
		}

		req := ArithRequest{9, 2}
		var res ArithResponse
		fmt.Println(time.Now(), "second call")
		err2 = conn2.Call("Arith.Multiply", req, &res)
		if err2 != nil {
			log.Println("arith error: ", err2)
		}
		log.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
	}()

	select {
	case <-time.After(timeout):
		log.Printf("[WARN] rpc call timeout\n")
		conn.Close()
		//reconnect

		ch <- struct{}{}
	}

	time.Sleep(time.Second * 120)

}
