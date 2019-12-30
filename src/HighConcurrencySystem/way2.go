//这并没有带来什么提升，我们用一个缓冲的队列替换了有缺陷的并发，也只是推迟了问题的产生时间而已。
//我们的同步处理器每次只向S3上传一个有效载荷（单线程），由于传入请求的速率远远大于单个处理器上传到S3的能力，
//因此我们的buffer channel迅速达到极限，队列已经阻塞并且无法再往里边添加作业。
package HighConcurrencySystem

import (
	"encoding/json"
	"io"
	"net/http"
)

var Queue chan Payload

func init() {
	Queue = make(chan Payload, MaxQueue)
}

func payloadHandler2(w http.ResponseWriter, r *http.Request) {
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
		Queue <- payload
	}

	w.WriteHeader(http.StatusOK)
}

func StartProcessor() {
	for {
		select {
		case job := <-Queue:
			job.UploadToS3() // <-- STILL NOT GOOD
		}
	}
}
