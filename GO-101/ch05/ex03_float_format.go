package main

import "fmt"

func FloatFormat() {
	var a = 123.12345
	var b = 3.14

	fmt.Printf("%06.3f, %06.3f\n", a, b) //min length(including decimal point): 6, decimal length: 3, filler:0
	fmt.Printf("%06.3g, %06.3g\n", a, b) //total printed number: 3
	fmt.Printf("%6.5g, %6.5g\n", a, b) //total printed number: 5
}