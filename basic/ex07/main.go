package main

import (
	"fmt"
)

// interfaces
type nameSayer interface {
	sayName() string
}

type nearby struct {
	somebody []nameSayer
}

func (n nearby) whoIsHere() {
	for _, s := range n.somebody {
		fmt.Println(s.sayName())
	}
}

type person struct {
	name string
}

func (p person) sayName() string {
	return p.name
}

type cat struct{}

func (p cat) sayName() string {
	return "...meow..."
}

func main() {
	place := nearby{
		somebody: []nameSayer{
			person{
				name: "Igor",
			},
			person{
				name: "Danilo",
			},
			cat{},
		},
	}
	place.whoIsHere()
}
