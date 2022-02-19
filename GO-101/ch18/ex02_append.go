package main

import "fmt"

func Append() {
	var slice1 = [] int {1, 2, 3}
	fmt.Println("slice1:", slice1)

	var slice2 = append(slice1, 4) //append() adds a new element to a slice and returns a new slice
	fmt.Println("slice2", slice2)
	fmt.Println("slice1 after: ", slice1)
}