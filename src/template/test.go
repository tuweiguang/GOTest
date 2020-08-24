package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", funcHello)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Printf("http server failed,err:%v\n", err)
		return
	}
}

func funcHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./test.tmpl")
	if err != nil {
		log.Printf("parse failed,err:%v\n", err)
	}

	_ = t.Execute(w, "twg")
}
