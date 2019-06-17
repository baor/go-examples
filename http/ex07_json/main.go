package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Example of http server with JSON unmarshalling and panics

// curl -X POST -d '{"data":"something"}' http://localhost:8080/test
type reqestData struct {
	Data string `json:"data"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
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

func main() {
	s := http.Server{Addr: ":8080"}
	http.HandleFunc("/test", testHandler)
	res := s.ListenAndServe()
	fmt.Println(res)
}
