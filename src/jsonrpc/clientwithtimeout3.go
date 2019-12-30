//跨语言RPC
package main

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
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

	conn, err = jsonrpc.Dial("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	var a, b int
	req := ArithRequest{9, 2}
	var res ArithResponse

	log.Println("first call")
	err = conn.Call("Arith.Mylock", a, &b) // 乘法运算
	if err != nil {
		log.Println("arith error: ", err)
	}

	log.Println("second call")
	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
	if err != nil {
		log.Println("arith error: ", err)
	}
	log.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	log.Println("third call")
	err = conn.Call("Arith.Myunlock", a, &b) // 乘法运算
	if err != nil {
		log.Println("arith error: ", err)
	}

}
