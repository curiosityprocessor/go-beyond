package main

import "fmt"

func MultiReturn() {
	result, isSuccess := Divide(7, 2)
	fmt.Println(result,isSuccess)
	result, isSuccess = Divide(3, 0)
	fmt.Println(result,isSuccess)
}

/*
	multiple return types encased in ( )
*/
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}