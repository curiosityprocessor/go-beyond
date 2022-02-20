package main

import "fmt"

func CopySlice() {
	slice1 := [] int {1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	fmt.Println("\n>>>copying values from slice1 to slice2")
	//copy by loop -> inefficient
	// for i, v := range slice1 {
	// 	slice2[i] = v
	// }

	//copy by function copy()
	copy(slice2, slice1) //copy slice1 values to slice2

	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	fmt.Println("\n>>>update slice1[2] to 99")
	slice1[2] = 99

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

}
