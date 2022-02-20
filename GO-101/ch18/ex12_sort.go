package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int {4,1,3,5,2} //unordered slice
	fmt.Println("\n>>> init slice:", slice)

	sort.Ints(slice)
	fmt.Println("\n>>> after sort slice:", slice)
	
}