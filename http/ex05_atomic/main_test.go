package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandlerPP(t *testing.T) {
	initIndex := int64(0)
	index = &initIndex
	N := 1000

	testFunc := func(t *testing.T) {
		t.Parallel()
		for i := 0; i < N; i++ {
			w1 := httptest.NewRecorder()
			r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
			testHandlerPP(w1, r1)

		}
		assert.True(t, atomic.LoadInt64(index) >= int64(N))
	}

	t.Run("testGroup", func(t *testing.T) {
		t.Run("test1", testFunc)
		t.Run("test2", testFunc)
		t.Run("test3", testFunc)
	})

	time.Sleep(1 * time.Second)

	assert.Equal(t, int64(3*N), atomic.LoadInt64(index))
}

func TestHandlerAdd(t *testing.T) {
	initIndex := int64(0)
	index = &initIndex
	N := 1000

	testFunc := func(t *testing.T) {
		t.Parallel()
		for i := 0; i < N; i++ {
			w1 := httptest.NewRecorder()
			r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
			testHandlerAdd(w1, r1)

		}
		assert.True(t, atomic.LoadInt64(index) >= int64(N))
	}

	t.Run("testGroup", func(t *testing.T) {
		t.Run("test1", testFunc)
		t.Run("test2", testFunc)
		t.Run("test3", testFunc)
	})

	time.Sleep(1 * time.Second)

	assert.Equal(t, int64(3*N), atomic.LoadInt64(index))
}
