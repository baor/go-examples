package main

import (
	"fmt"
	"net/http"
)

// Example of http server with asyncronus communitcation with channel

var indexCh chan int

func testHandler(w http.ResponseWriter, r *http.Request) {

	select {
	case index := <-indexCh:
		index++
		indexCh <- index
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Response from handler number %d", index)))
	default:
		w.WriteHeader(500)
	}
}

func main() {
	indexCh = make(chan int)
	indexCh <- 0
	s := http.Server{Addr: ":8080"}
	http.HandleFunc("/test", testHandler)
	res := s.ListenAndServe()
	fmt.Println(res)
}
