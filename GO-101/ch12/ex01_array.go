package main

import "fmt"

func Arrays() {
	var temperatures [5]float32 = [5]float32{-1.2, -3.0, 1.5, 6.2, 4.7}

	for i := 0; i < len(temperatures); i++ {
		fmt.Println(i, ": ", temperatures[i])
		//unformatted printing -> real numbers are printed with minimal decimal point i.e. 3.0 -> 3
	}

	arr := [...]int{90, 10, 25} //array declared with no explicit size [...]. size is set with initialized elements

	for i, v := range arr {
		fmt.Println(i, ": ", v)
	}
	
}