//跨语言RPC
package main
import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// 算数运算结构体
type Arith struct {
}

// 算数运算请求结构体
type ArithRequest struct {
	//A int
	//B int
	ParaMap map[string]int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

// 乘法运算方法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	if _,ok := req.ParaMap["A"]; !ok{
		return errors.New("No A of ParaMap")
	}
	if _,ok := req.ParaMap["B"]; !ok{
		return errors.New("No B of ParaMap")
	}
	res.Pro = req.ParaMap["A"] * req.ParaMap["B"]
	return nil
}

// 除法运算方法
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if _,ok := req.ParaMap["A"]; !ok{
		return errors.New("No A of ParaMap")
	}
	if _,ok := req.ParaMap["B"]; !ok{
		return errors.New("No B of ParaMap")
	}
	if req.ParaMap["B"] == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.ParaMap["A"] / req.ParaMap["B"]
	res.Rem = req.ParaMap["A"] % req.ParaMap["B"]
	return nil
}

func main() {
	rpc.Register(new(Arith)) // 注册rpc服务
	lis, err := net.Listen("tcp4", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", "start connection")
	for {
		conn, err := lis.Accept() // 接收客户端连接请求
		if err != nil {
			continue
		}
		go func(conn net.Conn) { // 并发处理客户端请求
			fmt.Fprintf(os.Stdout, "%s,my add:%v peer add %v\n", "new client in coming",conn.LocalAddr(),conn.RemoteAddr())
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

