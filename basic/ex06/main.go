package main

import (
	"fmt"
)

// goroutine context
type User struct {
	Name string
}

func main() {
	u := User{
		Name: "Travix",
	}
	for k := 0; k < 12; k++ {
		go func() {
			j := 0
			for i := 0; i < 10000000; i++ {
				j++
			}
		}()
	}

	go func() {
		fmt.Printf("Name: %s\n", u.Name)
	}()

	j := 0
	for i := 0; i < 10000000; i++ {
		j++
	}

	u.Name = "xivart"

	fmt.Printf("in the end: %+v\n", u)
}
