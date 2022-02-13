package main

import "fmt"

func main(){
	var x int32 = 3
	var y int32 = 2

	fmt.Println("x + y = ", x + y)
	fmt.Println("x - y = ", x - y)
	fmt.Println("x * y = ", x * y)
	fmt.Println("x / y = ", x / y)

	var p float32 = 3.14
	var q float32 = 7

	fmt.Println("p * q = ", p * q) //expected: 21.98 computed: 21.980001
	fmt.Println("p / q = ", p / q)

	// fmt.Println(x + p) //invalid operation: x + p (mismatched types int32 and float32)
	// fmt.Println(y * q) //invalid operation: y * q (mismatched types int32 and float32)
}