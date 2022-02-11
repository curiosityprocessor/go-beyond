package main

import "fmt"

func main() {
	var a int = 12
	var b float64 = 123456789.987654231
	var c string = "hello"

	fmt.Print("a:", a, "b", b)
	fmt.Println("c:", c)
	fmt.Printf("a: %d, b: %f, c: %s", a, b, c)
}