package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	indexWCh = make(chan int)
	indexRCh = make(chan int)
	go indexWriter()

	N := 1000

	testFunc := func(t *testing.T) {
		t.Parallel()
		for i := 0; i < N; i++ {
			w1 := httptest.NewRecorder()
			r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
			testHandler(w1, r1)
		}
	}

	t.Run("testGroup", func(t *testing.T) {
		t.Run("test1", testFunc)
		t.Run("test2", testFunc)
		t.Run("test3", testFunc)
	})

	close(indexWCh)
	time.Sleep(1 * time.Second)

	assert.Equal(t, 3*N, globalIndex)
}
