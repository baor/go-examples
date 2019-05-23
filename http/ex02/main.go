package main

import (
	"fmt"
	"net/http"
)

// Example of http server with simple handler

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler is called")
	w.WriteHeader(200)
	w.Write([]byte("Response from handler"))
}

func main() {
	s := http.Server{Addr: ":8080"}
	http.HandleFunc("/test", handler)
	fmt.Println(s.ListenAndServe())
}
