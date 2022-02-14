package main

import "fmt"

func main() {
	result, isSuccess := NamedDivide(7, 2)
	fmt.Println(result,isSuccess)
	result, isSuccess = NamedDivide(3, 0)
	fmt.Println(result,isSuccess)
}

/*
	return values can be named, to be used within the function scope
*/
func NamedDivide(a,b int) (result int, isSuccess bool) {
	if b == 0 {
		result = 0
		isSuccess = false
		return
	}
	result = a / b
	isSuccess = true
	return
}