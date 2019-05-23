package main

import "fmt"

// empty interface
func foo(x interface{}) {
	fmt.Println(x)
}

type User struct {
	Name string
}

func main() {
	u := User{Name: "Travix"}

	foo(5)
	foo("string value")
	foo(u)
	foo(&u)
}
