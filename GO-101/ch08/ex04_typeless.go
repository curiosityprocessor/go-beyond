package main

import "fmt"

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100 //no error: untyped constants are allocated a type when copied to a variable
	// var b int = FloatPI * 100 //error: cannot use FloatPI * 100 (type float64) as type int in assignment

	fmt.Println(a)
}