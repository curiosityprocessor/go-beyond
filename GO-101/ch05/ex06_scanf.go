package main

import "fmt"

func Scanf() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d \n", &a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}