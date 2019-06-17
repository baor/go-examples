package main

import (
	"fmt"
	"net/http"
)

// Example of http server with asyncronus communitcation, shared resource and race detection

var index int

func testHandler(w http.ResponseWriter, r *http.Request) {
	index++

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Response from handler number %d", index)))
}

func main() {
	s := http.Server{Addr: ":8080"}
	http.HandleFunc("/test", testHandler)
	res := s.ListenAndServe()
	fmt.Println(res)
}
