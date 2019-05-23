package main

import "fmt"

// ordering in map

type intMap map[int]bool

func main() {
	imap := intMap{}
	for i := 0; i < 10; i++ {
		imap[i] = true
	}

	i := 0
	for k := range imap {
		if k != i {
			fmt.Printf("Mismatch! %d != %d\n", k, i)
		}
		i++
	}
}
