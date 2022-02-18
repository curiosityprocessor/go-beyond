package main

import "fmt"

func Slices() {
	var slice []int //slice type can be declared without size

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	// slice[1] = 10 //panic: runtime error: index out of range [1] with length 0
	fmt.Println(slice)

	//slice vs array
	var array = [...]int{1, 2, 3} //declaring array with fixed size of 3 [1 2 3]
	fmt.Println(array)

	var slice1 = [] int {1, 2, 3} //declaring slice with size 3 [1 2 3]
	fmt.Println(slice1)

	var slice2 = [] int {1, 3:2, 5:3} //declaring slice with [1 0 0 2 0 3]
	fmt.Println(slice2)

	var slice3 = make([]int, 3) //declaring slice using make() with size 3
	fmt.Println(slice3)
}