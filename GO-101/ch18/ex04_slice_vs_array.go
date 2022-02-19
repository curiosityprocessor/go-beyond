package main

import (
	"fmt"
)

func SliceVsArray() {
	array := [5]int {1, 2, 3, 4, 5}
	slice := []int {1, 2, 3, 4, 5}

	changeArray(array)
	fmt.Println(array) //value of original array unchanged
	changeSlice(slice)
	fmt.Println(slice) //value that original slice points to changes
}

func changeArray(array2 [5]int) { //pass by value: values of array is copied to new array
	array2[2] = 999 //3rd element of array2 variable is changed
}

//slice is a struct that consists of a pointer to an array
//when copied, the pointer holds same value => can work like pass by reference
func changeSlice(slice2 []int) { //pass by value: value of slice **struct** is copied to new slice
	slice2[2] = 999 //2nd element of slice that slice2 point to is changed
}