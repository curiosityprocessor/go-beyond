package main

import "fmt"

func Constants() {
	const C int = 1
	// fmt.Println(&C) //not addressable - error: cannot take the address of C
	var b int = C * 10 //value accesible
	// C = 20 //not modifiable - error: cannot assign to C (declared const)

	fmt.Println(b)

}