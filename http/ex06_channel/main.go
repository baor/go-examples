package main

import (
	"fmt"
	"net/http"
)

// Example of http server with asyncronus communitcation with channel
// deadlock example

var globalIndex int
var indexWCh chan int
var indexRCh chan int

func indexWriter() {
	localIndex := 0
	for {
		index, ok := <-indexWCh
		if !ok {
			globalIndex = localIndex
			close(indexRCh)
			return
		}
		localIndex += index
		indexRCh <- localIndex
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	indexWCh <- 1
	index, ok := <-indexRCh
	if !ok {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Response from handler number %d", index)))

}

func main() {
	indexWCh = make(chan int)
	indexRCh = make(chan int)
	go indexWriter()
	s := http.Server{Addr: ":8080"}
	http.HandleFunc("/test", testHandler)
	res := s.ListenAndServe()
	fmt.Println(res)
}
