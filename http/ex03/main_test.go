package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handler01(t *testing.T) {
	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
	go testHandler(w1, r1)

	w2 := httptest.NewRecorder()
	r2, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
	go testHandler(w2, r2)

	assert.Equal(t, http.StatusOK, w1.Code)
}
