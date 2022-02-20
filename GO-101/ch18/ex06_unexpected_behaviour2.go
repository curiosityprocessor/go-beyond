package main

import "fmt"

func UnexpectedBehaviour2() {
	slice1 := []int  {1, 2, 3} //declare and init slice with len:3, cap:3
	slice2 := append(slice1, 4, 5) 
	//append() checks slice1 -> no empty space
	//creates and points to new array with double size and copies value from original array

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	fmt.Println("\n>>> update slice1[2] to 99")
	slice1[2] = 99

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	//update only applies to slice1 since slice1 and slice2 points to different array

	fmt.Println("\n>>> append 10 to slice1")
	slice1 = append(slice1, 10)
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	//append only applies slice1
	
}