package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// uncomment parallel to get race detection

func TestHandler(t *testing.T) {
	testFunc := func(t *testing.T) {
		//t.Parallel()
		for i := 0; i < 100; i++ {
			w1 := httptest.NewRecorder()
			r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
			testHandler(w1, r1)
		}
		assert.True(t, index >= 100)
	}

	t.Run("testGroup", func(t *testing.T) {
		t.Run("test1", testFunc)
		t.Run("test2", testFunc)
		t.Run("test3", testFunc)
	})

	time.Sleep(1 * time.Second)

	assert.Equal(t, 300, index)
}
