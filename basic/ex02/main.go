package main

import "fmt"

type UserName string
type UserAge int

type User struct {
	Name UserName
	Age  UserAge
}

type Users []User // Slice - Like list

type UniqueUsers map[User]bool // map - Like Dict

func main() {
	users := Users{} // var users = new Users();

	uAlice := User{
		Name: "Alice",
		Age:  20,
	}

	users = append(users, uAlice) // users.Add(new User(){...})

	users = append(users, User{
		Name: "Bob",
		Age:  30,
	})

	users = append(users, User{
		Name: "Bob",
		Age:  30,
	})

	fmt.Printf("\n users: %+v\n", users)

	unique := UniqueUsers{}
	for i, u := range users { // foreach (var i, u in users) {..}
		fmt.Printf("%d: %v\n", i, u)
		unique[u] = true
	}
	fmt.Printf("\n unique: %+v\n", unique)
}
