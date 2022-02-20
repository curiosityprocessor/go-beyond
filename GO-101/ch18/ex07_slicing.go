package main

import "fmt"

func Slicing() {
	//targetArray[startIndex:endIndex]
	//returns a slice by **slicing** a part of an array
	//startIndex: inclusive
	//endIndex: not inclusive

	array := [5]int {1, 2, 3, 4, 5}

	slice := array[1:2]
	fmt.Println("slice:", slice, len(slice), cap(slice))
	//slice has 1 element of 2. => len:1
	//cap is length of original array - start index of slice = 5 - 1 = 4
	//len:1 cap:4 empty: 3

	fmt.Println("\n>>>update [1] value in array")
	array[1] = 99
	fmt.Println("slice:", slice, len(slice), cap(slice))
	//value change in array applies to slice

	fmt.Println("\n>>>append 10 to slice")
	slice = append(slice, 10)
	//slice has len: 1, cap:4
	//slice[0] points to array[1]
	//slice calls to update slice[1] to 10
	//slice[1] points to array[2]
	//array[2] set to 10
	fmt.Println("slice:", slice, len(slice), cap(slice))
	fmt.Println("array:", array)
}
