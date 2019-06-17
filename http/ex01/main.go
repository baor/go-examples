package main

import (
	"net/http"
)

// Simple example of http server which does nothing
// https://golang.org/pkg/net/http/

func main() {
	s := http.Server{Addr: ":8080"}
	s.ListenAndServe()
}
