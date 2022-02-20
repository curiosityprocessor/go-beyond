package main

import "fmt"

func main() {
	slice := []int {1, 2, 3, 4, 5}
	idx := 2
	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice = slice[:len(slice) - 1]
	fmt.Println("slice:", slice, len(slice), cap(slice))
}

//manual implementation of removing an element
// func manualRemove(slice []int, idx int) {
// 	for i := idx + 1; i < len(slice); i++ {
// 		slice[i-1] = slice[i]
// 	}
// 	slice = slice[:len(slice) - 1]
// }