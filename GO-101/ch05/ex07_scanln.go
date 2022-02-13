package main

import "fmt"

func scanln() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

	n2, err2 := fmt.Scanln(&a, &b) 
	//if error exists on first scan, second scan will read from possible residual data on the input stream rather than new input
	
	if err2 != nil {
		fmt.Println(n2, err2)
	} else {
		fmt.Println(n2, a, b)
	}
}