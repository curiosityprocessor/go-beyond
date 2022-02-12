package main

import "fmt"

func integerFormat() {
	var a = 123
	var b = 1234
	var c = 1234567890

	fmt.Printf("%5d, %5d, %5d\n", a, b, c) //min length: 5, filler: (empty space)
	fmt.Printf("%05d, %05d, %05d\n", a, b, c) //min length: 5, filler: 0
	fmt.Printf("%-5d, %-5d, %-5d\n", a, b, c) //min length: 5, filler: , left-align
	fmt.Printf("%-05d, %-05d, %-05d\n", a, b, c) //min length: 5, filler: , left-align
	
}