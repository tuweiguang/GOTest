package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

//func main() {
//	mux := http.NewServeMux()
//
//	// Convert the timeHandler function to a HandlerFunc type
//	th := http.HandlerFunc(timeHandler)
//	// And add it to the ServeMux
//	mux.Handle("/time", th)
//
//	log.Println("Listening...")
//	http.ListenAndServe(":3000", mux)
//}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/time", timeHandler) /* 等同于：
	   th := http.HandlerFunc(timeHandler) //将函数timeHandler转换为HandlerFunc
	   mux.Handle("/time", th)             //将HandlerFunc注册到mux
	*/

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
