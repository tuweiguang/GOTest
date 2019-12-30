package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //创建空的 ServeMux

	rh := http.RedirectHandler("http://example.org", 307) //重定向处理器
	mux.Handle("/foo", rh)                                //将处理器注册到新创建的 ServeMux,所以它在 URL 路径/foo 上收到所有的请求都交给这个处理器

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux) //通过mux创建了一个新的服务器,然后通过server.ListenAndServe()函数监听所有请求
	// mux :ServeMux 实现Handler接口，所以传mux合法
}
