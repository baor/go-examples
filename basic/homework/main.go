package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type diceThrower interface {
	throw() int
}

type nameSayer interface {
	sayName() string
}

type gamePlayer interface {
	nameSayer
	diceThrower
}

type players []gamePlayer

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

func (g *game) play() {
	for i := range g.players {
		p := g.players[i]
		name := p.sayName()
		score := p.throw()
		g.scores = append(g.scores, scoreLine{playerName: name, score: score})

		fmt.Printf("%s scores %d\n", name, score)
	}
	sort.Sort(g.scores)

	//fmt.Printf("scores after sorting: %+v\n", g.scores)
	g.winnerName = g.scores[len(g.scores)-1].playerName
}

type fairPlayer struct {
	name string
}

func (p fairPlayer) throw() int {
	return rand.Intn(6) + 1
}

func (p fairPlayer) sayName() string {
	return p.name
}

type cheater struct {
	name string
}

func (c cheater) throw() int {
	return rand.Intn(11) + 2
}

func (c cheater) sayName() string {
	return c.name
}

func main() {
	rand.Seed(time.Now().UnixNano())
	simpleGame := game{
		players: players{
			fairPlayer{
				name: "Player1",
			},
			fairPlayer{
				name: "Player2",
			},
			cheater{
				name: "Cheater",
			},
		},
	}
	simpleGame.play()

	fmt.Printf("Winner is %+v !\n", simpleGame.winnerName)
}
