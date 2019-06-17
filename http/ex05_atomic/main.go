package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

// Example of http server with asyncronus communitcation with atomic.

var index *int64

func testHandlerPP(w http.ResponseWriter, r *http.Request) {

	value := atomic.LoadInt64(index)
	value++
	atomic.StoreInt64(index, value)

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Response from handler number %d", index)))
}

func testHandlerAdd(w http.ResponseWriter, r *http.Request) {

	value := atomic.AddInt64(index, 1)
	atomic.StoreInt64(index, value)

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Response from handler number %d", index)))
}

func main() {
	initIndex := int64(0)
	index = &initIndex
	s := http.Server{Addr: ":8080"}
	http.HandleFunc("/test", testHandlerPP)
	res := s.ListenAndServe()
	fmt.Println(res)
}
