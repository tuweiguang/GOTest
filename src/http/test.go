package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//http请求
func httpHandle(method, urlVal, data string) {
	client := &http.Client{}
	var req *http.Request

	if data == "" {
		urlArr := strings.Split(urlVal, "?")
		if len(urlArr) == 2 {
			urlVal = urlArr[0] + "?" + getParseParam(urlArr[1])
		}
		req, _ = http.NewRequest(method, urlVal, nil)
	} else {
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}

	//添加cookie，key为X-Xsrftoken，value为df41ba54db5011e89861002324e63af81
	//可以添加多个cookie
	cookie1 := &http.Cookie{Name: "X-Xsrftoken", Value: "df41ba54db5011e89861002324e63af81", HttpOnly: true}
	req.AddCookie(cookie1)

	//添加header，key为X-Xsrftoken，value为b6d695bbdcd111e8b681002324e63af81
	req.Header.Add("X-Xsrftoken", "b6d695bbdcd111e8b681002324e63af81")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

//将get请求的参数进行转义
func getParseParam(param string) string {
	return url.PathEscape(param)
}

//测试
func main() {
	httpHandle("GET", "http://www.baidu.com", "")
}
