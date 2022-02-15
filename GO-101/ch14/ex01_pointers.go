package main

import "fmt"

func Pointers() {
	var a int = 104
	var p *int = &a

	fmt.Printf("value of a: %d\n", a)
	fmt.Printf("pointer value: %p\n", p)
	fmt.Printf("value at pointer: %d\n", *p)

	*p = 96
	fmt.Println("===== value changed using pointer")
	fmt.Printf("new value of a: %d\n", a)

}