package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_handler01(t *testing.T) {

	indexCh = make(chan int, 1)
	indexCh <- 0

	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
	go testHandler(w1, r1)

	w2 := httptest.NewRecorder()
	r2, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
	go testHandler(w2, r2)

	time.Sleep(100)
	select {
	case value := <-indexCh:
		assert.Equal(t, 2, value)
	default:
		assert.Equal(t, "B", "A")
	}
}
