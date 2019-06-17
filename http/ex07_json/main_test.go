package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handler01(t *testing.T) {

	someData := reqestData{Data: "some data inside"}
	reqBody, _ := json.Marshal(someData)

	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(reqBody))
	testHandler(w1, r1)

	assert.Equal(t, http.StatusOK, w1.Code)

	respBody, _ := ioutil.ReadAll(w1.Body)
	assert.Equal(t, "Data from handler 'some data inside'", string(respBody))
}

func Test_handler02(t *testing.T) {

	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("invalid")))
	testHandler(w1, r1)

	assert.Equal(t, http.StatusOK, w1.Code)

	respBody, _ := ioutil.ReadAll(w1.Body)
	assert.Equal(t, "Data from handler 'some data inside'", string(respBody))
}
