package main

import (
	"fmt"
	"net/http"
)

// Simple example of http server which does nothing
// https://golang.org/pkg/net/http/

func main() {
	s := http.Server{Addr: ":8080"}
	fmt.Println(s.ListenAndServe())
}
