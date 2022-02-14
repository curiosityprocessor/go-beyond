package main

import "fmt" //imports fmt package for standard I/O operations

func FmtPackage() {
	var a int = 12
	var b float64 = 123456789.987654231
	var c string = "hello"

	fmt.Print("a:", a, "b", b) //prints
	fmt.Println("c:", c) //prints + line change
	fmt.Printf("a: %d, b: %f, c: %s", a, b, c) //print with format
	fmt.Print("end")
}