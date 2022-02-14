package main

import "fmt"

/*
	unlike other languages, switch statements in GO does not require break keyword
	first true case block is run, then the switch code block is exited
*/
func break() {
	a := 1

	switch {
		case a == 1: //first occurencce of true statement
			fmt.Println("a is 1") //code block is run, then implicit break, switch block exited
		case a < 2: //even though true, this block is not run
			fmt.Println("a is 2")
			break //break is optional
		default:
			fmt.Println("no idea")
	}
}