package main

import "net/rpc"
import (
	. "../protocol"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type Calculator struct{}

var (
	_DATA       *Calculator
	_CAN_CANCEL chan bool
)

func main() {
	runRpcServer()
}

func init() {
	_DATA = new(Calculator)
	_CAN_CANCEL = make(chan bool)
}

func runRpcServer() {
	//rpc包里面定义了个DefaultServer，缺省的Register和HandleHTTP均是对DefaultServer作的操作，如果想定制新的Server，就自己写
	rpc.Register(_DATA)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:2311")
	if e != nil {
		fmt.Errorf("Create Listener Error %s", e.Error())
	}

	go http.Serve(l, nil)
	//阻塞主进程，等待客户端输入
	<-_CAN_CANCEL
}

//输出方法的格式要求：func (t *T) MethodName(argType T1, replyType *T2) error
func (*Calculator) Addition(param *Param, reply *int32) error {
	*reply = param.A + param.B

	return nil
}

func (*Calculator) Subtraction(param *Param, reply *int32) error {
	*reply = param.A - param.B
	return nil
}

func (*Calculator) Multiplication(param *Param, reply *int32) error {
	*reply = param.A * param.B

	return nil
}

func (*Calculator) Division(param *Param, reply *int32) error {
	if 0 == param.B {
		return errors.New("divide by zero")
	}
	*reply = param.A / param.B

	return nil
}
