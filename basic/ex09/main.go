package main

import "fmt"

func foo1() {
	defer fmt.Println("foo1-1")
	fmt.Println("foo1-2")
}

func foo2() {
	defer fmt.Println("foo2-1")
	fmt.Println("foo2-2")
}

func main() {
	defer foo2()
	defer fmt.Println("main-1")
	defer fmt.Println("main-2")
	defer fmt.Println("main-3")
	fmt.Println("main-4")
	foo1()
	defer fmt.Println("main-5")
}
