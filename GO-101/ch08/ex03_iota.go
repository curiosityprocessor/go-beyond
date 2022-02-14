package main

import "fmt"
const (
	Red int = iota
	Green int = iota
	Blue int = iota
)

const (
	C1 int = iota + 1
	C2
	C3
)

func IotaConstants() {
	fmt.Println(Red)
	fmt.Println(Green)
	fmt.Println(Blue)
	fmt.Println("===========")
	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(C3)
}