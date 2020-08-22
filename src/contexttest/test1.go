package main

import (
	"context"
	"log"
	"os"
	"time"
)

var MyPrint *log.Logger

func InitLog() {
	MyPrint = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Cdd(ctx context.Context) int {
	log.Println(ctx.Value("NLJB"))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		log.Println("====1====")
		return -3
	}
}

func Bdd(ctx context.Context) int {
	log.Println(ctx.Value("HELLO"))
	log.Println(ctx.Value("WROLD"))
	ctx = context.WithValue(ctx, "NLJB", "NULIJIABEI")
	go log.Println(Cdd(ctx))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		log.Println("====2====")
		return -2
	}
}
func Add(ctx context.Context) int {
	ctx = context.WithValue(ctx, "HELLO", "WROLD")
	ctx = context.WithValue(ctx, "WROLD", "HELLO")
	go log.Println(Bdd(ctx))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		log.Println("====3====")
		return -1
	}
}
func main() {
	InitLog()

	// 自动取消(定时取消)
	{
		timeout := 3 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		log.Println(Add(ctx))
	}

	// 手动取消
	//  {
	//      ctx, cancel := context.WithCancel(context.Background())
	//      go func() {
	//          time.Sleep(2 * time.Second)
	//          cancel() // 在调用处主动取消
	//      }()
	//      fmt.Println(Add(ctx))
	//  }
	time.Sleep(5 * time.Second)
}
