package main

import "fmt"

func main() {
	arr1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("arr1 is")
	for  i, v := range arr1 {
		println(i, v)
	}

	arr2 := [5]int{6, 7, 8, 9, 0}
	fmt.Println("\narr2 is")
	for  i, v := range arr2 {
		println(i, v)
	}

	arr2 = arr1 //copying array
	
	fmt.Println("\narr2 is now")
	for  i, v := range arr2 {
		println(i, v)
	}


}