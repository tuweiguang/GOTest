package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://192.168.18.165:9999/api/v2/write?org=mercury&bucket=server"

	payload := strings.NewReader("tuweiguang,Ip_address=192.168.3.141 myfield=77")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", " Token j2n_nQHTrdFPQ9RU2xDyq32cFbQ8KFbegIVHXOVlEDLAQIfcJxqg1f8523yiyxan006BCXUB0AFPGt02cUJwOA==")

	res, err1 := http.DefaultClient.Do(req)
	if res == nil && err1 != nil {
		return
	}
	defer res.Body.Close()
	fmt.Printf("StatusCode:%v err:%v\n", res, err1)
	if res.StatusCode >= 200 && res.StatusCode < 400 && err1 == nil {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(res)
		fmt.Println(string(body))
	}

}
