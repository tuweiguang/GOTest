package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 调用cancelFunc或者下游goroutine运行超过5s，将关闭下游goroutine
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	ctx = context.WithValue(ctx, "Test", "123456")
	//defer cancelFunc()

	go func() {
		time.Sleep(time.Second * 2)
		cancelFunc()
	}()
	if t, ok := ctx.Deadline(); ok {
		fmt.Println(time.Now())
		fmt.Println(t.String())
	}
	go func(ctx context.Context) {
		fmt.Println(time.Now(), ctx.Value("Test"))
		for {
			select {
			case <-ctx.Done():
				fmt.Println(time.Now(), ctx.Err())
				return
			default:
				continue
			}
		}
	}(ctx)

	time.Sleep(time.Second * 10)
}
