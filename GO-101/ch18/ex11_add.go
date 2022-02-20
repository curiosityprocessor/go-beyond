package main

import "fmt"

func Add() {
	slice := [] int {1, 2, 3, 4, 5}
	fmt.Println("\n>>> init\nslice:", slice, len(slice), cap(slice))

	idx := 2 //index to add element to
	value := 99
	fmt.Printf("\n>>> updating slice at index [%d] to %d\n", idx, value)

	//1. manual implementation
	// slice = manualAdd(slice, idx, value)
	
	//2. using append
	slice = addByAppend(slice, idx, value)

	fmt.Println("\n>>> after update\nslice:", slice, len(slice), cap(slice))

}

func manualAdd(slice []int, idx int, value int) ([] int) {
	fmt.Println("\n===== 1.manual implementation:")

	//add new element to end
	slice = append(slice, 0) 
	fmt.Println("\n>>> after append 0\nslice:", slice, len(slice), cap(slice))

	//loop from end - 1 to idx, set n value to n + 1
	for i := len(slice) - 1 - 1; i >= idx ; i-- {		
		slice[i+1] = slice[i]
	}
	fmt.Println("\n>>> after loop\nslice:", slice, len(slice), cap(slice))

	//update target idx value
	slice[idx] = value
	return slice
}

func addByAppend(slice []int, idx int, value int) ([] int) {
	fmt.Println("\n===== 2.using append():")

	//add new element to end
	slice = append(slice, 0)
	fmt.Println("\n>>> after append 0\nslice:", slice, len(slice), cap(slice))

	//use copy() to move values
	copy(slice[idx+1:], slice[idx:])
	fmt.Println("\n>>> after copy\nslice:", slice, len(slice), cap(slice))

	//update target idx value
	slice[idx] = value
	return slice
}