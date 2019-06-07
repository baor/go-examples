package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// func Test_handler01(t *testing.T) {
// 	simpleGame = game{
// 		players: players{
// 			"Player1": fairPlayer{},
// 			"Player2": fairPlayer{},
// 			"Cheater": cheater{},
// 		},
// 	}

// 	w1 := httptest.NewRecorder()
// 	r1, _ := http.NewRequest("POST", "/play", bytes.NewBuffer([]byte(`{"name": "Player1"}`)))
// 	go handlerPlay(w1, r1)

// 	w2 := httptest.NewRecorder()
// 	r2, _ := http.NewRequest("POST", "/play", bytes.NewBuffer([]byte(`{"name": "Player2"}`)))
// 	go handlerPlay(w2, r2)

// 	w3 := httptest.NewRecorder()
// 	r3, _ := http.NewRequest("POST", "/play", bytes.NewBuffer([]byte(`{"name": "Cheater"}`)))
// 	go handlerPlay(w3, r3)
// 	time.Sleep(500)

// 	assert.Equal(t, http.StatusOK, w1.Code)
// 	assert.Equal(t, http.StatusOK, w2.Code)
// 	assert.Equal(t, http.StatusOK, w3.Code)
// }

func Test_handler01(t *testing.T) {
	simpleGame = game{
		players: players{
			"Player1": fairPlayer{},
			"Player2": fairPlayer{},
			"Cheater": cheater{},
		},
	}

	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", fmt.Sprintf("/play?name=%s", "Player1"), nil)
	go handlerPlay(w1, r1)

	w2 := httptest.NewRecorder()
	r2, _ := http.NewRequest("GET", fmt.Sprintf("/play?name=%s", "Player2"), nil)
	go handlerPlay(w2, r2)

	w3 := httptest.NewRecorder()
	r3, _ := http.NewRequest("GET", fmt.Sprintf("/play?name=%s", "Cheater"), nil)
	go handlerPlay(w3, r3)
	time.Sleep(500)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, http.StatusOK, w3.Code)
}
