package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", funcHello)
	http.HandleFunc("/test2", test2)
	http.HandleFunc("/defineFunc", defineFunc)
	http.HandleFunc("/login", login)
	http.HandleFunc("/main", base)
	http.HandleFunc("/main/user", user)
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

func test2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./test2.tmpl")
	if err != nil {
		log.Printf("parse failed,err:%v\n", err)
	}

	s := []int{
		1, 2, 3, 4, 5,
	}
	_ = t.Execute(w, s)
}

func defineFunc(w http.ResponseWriter, r *http.Request) {

	// 要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	f := func(arg string) (string, error) {
		return "hello, " + arg, nil
	}

	// 注意new里面的模板名字和解析模板文件名字一样
	t, err := template.New("func.tmpl").Funcs(template.FuncMap{"myfunc": f}).ParseFiles("./func.tmpl")
	if err != nil {
		log.Printf("parse failed, err:%v\n", err)
		return
	}

	_ = t.ExecuteTemplate(w, "func.tmpl", "twg")
}

func login(w http.ResponseWriter, r *http.Request) {
	var password string
	var username string

	username = r.FormValue("username")
	password = r.FormValue("password")
	log.Println("===> ", username, password)

	if username == "admin" && password == "123456" {
		http.Redirect(w, r, "/main", 302)
	}

	t, err := template.ParseFiles("./login.tmpl")
	if err != nil {
		log.Printf("parse failed,err:%v\n", err)
	}

	_ = t.Execute(w, "")
}

func base(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./main.tmpl")
	if err != nil {
		log.Printf("parse failed,err:%v\n", err)
	}

	_ = t.Execute(w, "")
}

func user(w http.ResponseWriter, r *http.Request) {
	// 继承基础tmpl要放第一个
	t, err := template.ParseFiles("./main.tmpl", "./user.tmpl")
	if err != nil {
		log.Printf("parse failed,err:%v\n", err)
	}

	_ = t.ExecuteTemplate(w, "user.tmpl", map[string]interface{}{
		"playerId": 123456,
		"uid":      123456,
		"name":     "twg",
		"age":      18,
	})
}
