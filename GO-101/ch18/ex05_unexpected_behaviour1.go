package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5) //len:3, cap:5
	//slice1 has 2 empty spaces left (cap - len = 2)

	slice2 := append(slice1, 4, 5) //appends 4, 5 to slice1

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[2] = 99 //set 3rd value of array that slice1 points to to 99

	fmt.Println("\n>>> set slice1[2] to 99")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	//both slice1, slice2 show value changes(points to same array)

	slice1 = append(slice1, 6) //append 6 to slice1
	//slice1 still has 2 empty spaces left (cap - len = 2)

	fmt.Println("\n>>> after append 6 to slice1")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	//slice1[3] sets to 6 by append
	//slice2[3] was 4, is now 6 


}