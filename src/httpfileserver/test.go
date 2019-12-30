package main

import (
	"fmt"
	"net/http"
)

/* 1 */
//func main()  {
//	// 访问127.0.0.1:8080
//	http.Handle("/",http.StripPrefix("/",http.FileServer(http.Dir("test"))))
//
//	err := http.ListenAndServe(":8080",nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

/* 2 */
//func main()  {
//	// 访问127.0.0.1:8080/test
//	http.Handle("/test/",http.StripPrefix("/test/",http.FileServer(http.Dir("test"))))
//
//	err := http.ListenAndServe(":8080",nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

/* 3 */
//func main()  {
//	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("test")))
//	if err != nil {
//		fmt.Println(err)
//	}
//}

/* 4 */
//func main()  {
//	handle := http.FileServer(http.Dir("test"))
//	http.Handle("/",handle )
//
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

/* 5 */
//var fileServer http.Handler
//func helloHandler(rw http.ResponseWriter, req *http.Request) {
//	url := req.Method + " " + req.URL.Path
//	fmt.Println(url)
//	if req.URL.RawQuery != "" {
//		url += "?" + req.URL.RawQuery
//	}
//	fileServer.ServeHTTP(rw,req)
//}
//
//func main()  {
//	fileServer = http.FileServer(http.Dir("test"))
//	http.HandleFunc("/", helloHandler)    // HandlerFunc实现了普通函数转Handler接口
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

/* 6 */
var fileServer http.Handler

func hello1Handler(rw http.ResponseWriter, req *http.Request) {
	url := req.Method + " " + req.URL.Path
	fmt.Println(url)
	if req.URL.RawQuery != "" {
		url += "?" + req.URL.RawQuery
	}
	fileServer.ServeHTTP(rw, req)
}

func hello2Handler(rw http.ResponseWriter, req *http.Request) {
	url := req.Method + " " + req.URL.Path
	fmt.Println(url)
	if req.URL.RawQuery != "" {
		url += "?" + req.URL.RawQuery
	}

	rw.Write([]byte("hello"))
}

func main() {
	fileServer = http.FileServer(http.Dir("test"))
	http.HandleFunc("/", hello1Handler)
	//可以处理动态事务         因为路径的不同返回的结果也不相同
	http.HandleFunc("/test/", hello2Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

/*
   rw http.ResponseWriter  ==> 返回要显示的数据
   req *http.Request       ==> 接受的数据
*/

//参考网址 http://www.cnblogs.com/shockerli/archive/2018/08/21/golang-pkg-http-file-server.html

//http.Dir("test") 是将当前路径的test文件或目录 换成文件系统
//http.FileServer(http.Dir("test")) 如果访问路径是目录，则列出目录内容；如果访问路径是文件，则显示文件内容
//http.FileServer() 方法返回的是 fileHandler 实例，而 fileHandler 结构体实现了 Handler 接口的方法 ServeHTTP()。ServeHTTP 方法内的核心是 serveFile() 方法。
//http.HandleFunc() 或 http.Handle() 可以实现带路由前缀的文件服务。比如：127.0.0.1:8080 或 127.0.0.1:8080/test   如果没有这两个函数，则这两者的地址都可以输入
//http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))  FileServer 已经明确静态文件的根目录在"/tmp"，但是我们希望URL以"/tmpfiles/"开头。
// 如果有人请求"/tempfiles/example.txt"，我们希望服务器能将文件发送给他。为了达到这个目的，我们必须从URL中过滤掉"/tmpfiles", 而剩下的路径是相对于根目录"/tmp"的相对路径。
// 如果我们按照如上做法，将会得到如下结果： /tmp/example.txt
