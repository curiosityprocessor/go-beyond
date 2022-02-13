package main

import "fmt"

func scan() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b) //parameters: the address of variables (designated with &)
	if err != nil {
		fmt.Println(n, err) //when error
	} else {
		fmt.Println(n, a, b) //when success
	}
}