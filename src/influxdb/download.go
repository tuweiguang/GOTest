package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://192.168.18.165:9999/api/v2/query?org=mercury"

	payload := strings.NewReader("from(bucket: \"ue4\")|> range(start:-1000h)")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", " Token w96Wr3ZJa2x0ttb3mFzR3vjJJyZy71FWQnyU1FSqSNhs6gRebvuG3SDRHdScfk26mQaZGJ_U3KIBoTGFp7H0gQ==")
	req.Header.Add("accept", "application/csv")
	req.Header.Add("content-type", "application/vnd.flux")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

//curl "http://192.168.18.165:9999/api/v2/query?org=mercury" -XPOST -sS -H "Authorization: Token j2n_nQHTrdFPQ9RU2xDyq32cFbQ8KFbegIVHXOVlEDLAQIfcJxqg1f8523yiyxan006BCXUB0AFPGt02cUJwOA==" -H "accept:application/csv" -H "content-type:application/vnd.flux" --data 'from(bucket: "test")|> range(start: v.timeRangeStart, stop: v.timeRangeStop)|> filter(fn: (r) => r._measurement == "stat")'
