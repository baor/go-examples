package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handler01(t *testing.T) {

	go runOkSrv()

	someData := reqestData{Data: "some data inside"}
	reqBody, _ := json.Marshal(someData)

	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(reqBody))
	testHandler(w1, r1)

	assert.Equal(t, http.StatusOK, w1.Code)

	respBody, _ := ioutil.ReadAll(w1.Body)
	assert.Equal(t, "Data from handler 'some data inside'", string(respBody))
}

func Benchmark_handler01(b *testing.B) {
	go runOkSrv()

	someData := reqestData{Data: "some data inside"}
	reqBody, _ := json.Marshal(someData)

	b.ResetTimer()
	printMemUsage()
	for i := 0; i < b.N; i++ {
		w1 := httptest.NewRecorder()
		r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(reqBody))
		testHandler(w1, r1)
	}
	printMemUsage()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats

	// Alloc is bytes of allocated heap objects.
	fmt.Printf("\nAlloc = %v MiB", bToMb(m.Alloc))
	// TotalAlloc is cumulative bytes allocated for heap objects.
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	// Sys is the total bytes of memory obtained from the OS.
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	// HeapSys is bytes of heap memory obtained from the OS.
	fmt.Printf("\tHeapSys = %v MiB", bToMb(m.HeapSys))
	fmt.Println()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
