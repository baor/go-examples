package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Example of http server with body leak

type reqestData struct {
	Data string `json:"data"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// leak is not here
	//defer r.Body.Close()

	req, err := http.NewRequest("GET", "http://localhost:8081/", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	// leak is here
	//defer resp.Body.Close()

	if resp == nil {
		panic("resp is nil")
	}

	if err != nil {
		panic(err)
	}

	rd := reqestData{}
	if err := json.Unmarshal(body, &rd); err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Data from handler '%s'", rd.Data)))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Somedata")))
}

func runOkSrv() {
	mux2 := http.NewServeMux()
	svr2 := http.Server{
		ReadTimeout:       time.Second * 2,
		ReadHeaderTimeout: time.Second * 2,
		WriteTimeout:      time.Second * 2,
		Addr:              ":8081",
		Handler:           mux2,
	}
	mux2.HandleFunc("/ok", testHandler)

	svr2.ListenAndServe()
}

func main() {
	mux := http.NewServeMux()
	svr := http.Server{
		ReadTimeout:       time.Second * 2,
		ReadHeaderTimeout: time.Second * 2,
		WriteTimeout:      time.Second * 2,
		Addr:              ":8080",
		Handler:           mux,
	}
	mux.HandleFunc("/test", testHandler)
	go runOkSrv()
	res := svr.ListenAndServe()
	fmt.Println(res)
}
