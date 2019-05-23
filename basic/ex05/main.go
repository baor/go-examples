package main

import "fmt"

// extensions and pointers

type User struct {
	Name string
}

func (u *User) pointer() { // foo (this x)
	fmt.Printf("p.u.address: %p\n", u)
	u.Name = "Poiner"
}
func (u User) copy() {
	fmt.Printf("c.u.address: %p\n", &u)
	u.Name = "copy"
}

func main() {
	u := User{
		Name: "Travix",
	}
	fmt.Printf("original address: %p\n", &u)
	fmt.Printf("before copy: %v\n", u)
	u.copy()
	fmt.Printf("after copy: %v\n", u)
	fmt.Printf("before pointer: %v\n", u)
	u.pointer()
	fmt.Printf("after pointer: %v\n", u)

}
