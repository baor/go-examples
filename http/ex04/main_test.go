package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handler01(t *testing.T) {

	initIndex := int64(0)
	index = &initIndex

	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
	go testHandler(w1, r1)

	w2 := httptest.NewRecorder()
	r2, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
	go testHandler(w2, r2)

	//time.Sleep(100)
	value := atomic.LoadInt64(index)
	assert.Equal(t, int64(2), value)
}
