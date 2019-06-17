package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// uncomment go to get race or t.Parallel() to see error
func testAction(n int) {
	for i := 0; i < n; i++ {
		w1 := httptest.NewRecorder()
		r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("")))
		/*go*/ testHandler(w1, r1)
	}
	time.Sleep(100)
}

func Test_handler01(t *testing.T) {
	index = 0
	//t.Parallel()
	testAction(100)
	assert.Equal(t, 100, index)
}

func Test_handler02(t *testing.T) {
	index = 0
	//t.Parallel()
	testAction(100)
	assert.Equal(t, 100, index)
}
