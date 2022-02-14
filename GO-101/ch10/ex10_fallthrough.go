package main

import "fmt"

/*
	other language break = go lang implicit break
	other language no explicit break = go lang fallthrough
*/
func main() {
	a := 3

	switch {
		case a == 1:
			fmt.Println("a is 1")
		case a < 2:
			fmt.Println("a is 2")
		case a == 3:
			fmt.Println("a is 3")
			fallthrough //fallthrough cancels break and runs the case block below
		case a == 4:
			fmt.Println("a is 4?")
		
	}
}