//最简单粗暴的方式，每个请求一个goroutine，但是对于高并发的请求，系统很快就会崩溃
package HighConcurrencySystem

import (
	"encoding/json"
	"io"
	"net/http"
)

//方法一
func payloadHandler1(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read the body into a string for json decoding
	var content = &PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range content.Payloads {
		go payload.UploadToS3() // <----- DON'T DO THIS
	}

	w.WriteHeader(http.StatusOK)
}
