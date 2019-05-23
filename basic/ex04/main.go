package main

import "fmt"

// array vs slice

func foo(s []int, a [2]int) {
	s[0] = 0
	a[0] = 0
}

func main() {
	s := []int{1, 2}
	a := [2]int{1, 2}
	fmt.Printf("before: s: %v, a: %v\n", s, a)
	foo(s, a)
	fmt.Printf("after: s: %v, a: %v\n", s, a)
}
