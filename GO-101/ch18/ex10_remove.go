package main

import "fmt"

func Remove() {
	slice := []int {1, 2, 3, 4, 5}
	idx := 2
	fmt.Println("\ninit\nslice:", slice, len(slice), cap(slice))

	// slice = manualRemove(slice, idx)

	slice = append(slice[:idx], slice[idx+1:]...) //remove using append
	//slice[:idx] = [1, 2]
	//slice[idx+1:] = [4, 5]

	fmt.Println("\nafter remove\nslice:", slice, len(slice), cap(slice))
}

//manual implementation of removing an element
func manualRemove(slice []int, idx int) ([]int) {
	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	return slice[:len(slice) - 1]
}