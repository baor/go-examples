package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

type diceThrower interface {
	throw() int
}

type gamePlayer interface {
	diceThrower
}

type players map[string]gamePlayer

type scoreLine struct {
	playerName string
	score      int
}

type scores []scoreLine

func (s scores) Len() int           { return len(s) }
func (s scores) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s scores) Less(i, j int) bool { return s[i].score < s[j].score }

type game struct {
	players    players
	scores     scores
	winnerName string
}

func (g *game) OneTurn(name string) {
	p, ok := g.players[name]
	if !ok {
		fmt.Printf("Wrong user name %s\n", name)
		return
	}
	score := p.throw()
	g.scores = append(g.scores, scoreLine{playerName: name, score: score})
	fmt.Printf("%s scores %d\n", name, score)
}

func (g *game) GetWinner() {
	if len(g.scores) < len(g.players) {
		fmt.Printf("Game is not over yet\n")
		return
	}

	sort.Sort(g.scores)
	g.winnerName = g.scores[len(g.scores)-1].playerName
	fmt.Printf("The winner is %s!!!\n", g.winnerName)
}

type fairPlayer struct {
	name string
}

func (p fairPlayer) throw() int {
	return rand.Intn(6) + 1
}

type cheater struct {
	name string
}

func (c cheater) throw() int {
	return rand.Intn(11) + 2
}

// type playRequest struct {
// 	Name string `json:"name"`
// }

var simpleGame game

func handlerPlay(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	// defer r.Body.Close()
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// req := playRequest{}
	// if err := json.Unmarshal(body, &req); err != nil {
	// 	panic(err)
	// }

	fmt.Printf("User %s does a turn\n", name)
	simpleGame.OneTurn(name)
	simpleGame.GetWinner()

	for {
		if len(simpleGame.winnerName) > 0 {
			break
		}
		time.Sleep(100)
	}

	w.WriteHeader(200)
	w.Write([]byte(simpleGame.winnerName))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	simpleGame = game{
		players: players{
			"Player1": fairPlayer{},
			"Player2": fairPlayer{},
			"Cheater": cheater{},
		},
	}

	s := http.Server{Addr: ":8080"}
	http.HandleFunc("/play", handlerPlay)
	fmt.Println(s.ListenAndServe())
}
