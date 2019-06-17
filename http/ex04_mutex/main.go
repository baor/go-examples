package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Example of http server with asyncronus communitcation with mutex.

var mutex *sync.Mutex
var index int

func testHandler(w http.ResponseWriter, r *http.Request) {

	mutex.Lock()
	index++
	mutex.Unlock()

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Response from handler number %d", index)))
}

func main() {
	mutex = &sync.Mutex{}
	s := http.Server{Addr: ":8080"}
	http.HandleFunc("/test1", testHandler)
	res := s.ListenAndServe()
	fmt.Println(res)
}
