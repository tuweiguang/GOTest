package

import (
	"net"
	"net/rpc"
	"log"
	"sync"
	"strings"
	"net/rpc/jsonrpc"
)

type HelloService struct {
	ip string
}

func (p *HelloService) Login(request string, reply *string) error {

	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {

	return nil
}

var mutex  sync.Mutex

func main() {
	server := rpc.NewServer()
	serverHandler := new(HelloService)
	server.Register(serverHandler)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go func() {
			mutex.Lock() //使用锁的话，高并发锁竞争时间增加
			serverHandler.ip = strings.Split(conn.RemoteAddr().String(),":")[0]
			server.ServeCodec(jsonrpc.NewServerCodec(conn))
			mutex.Unlock()
		} ()
	}
}

