package main

import "fmt"

func variableConversion() {
	/*
		[type conversions]
		Go is a [very] strongly type language
		does not offer any implicit/automated type conversions
		ex) 
			int8 -> int 16 (x)
			float64 -> int (x)

	*/

	a := 4	//inferred type is int
	var b float64 = 4.5

	// var c int = b //error: cannot use b (type float64) as type int in assignment

	// d := a * b //error: invalid operation: a * b (mismatched types int and float64)
	
	var e int16 = 3456 
	// var f int8 = e //error: cannot use g (type int16) as type int8 in assignment

	//using type conversions
	var c int = int(b) //convert from float64 -> int
	d := float64(a) * b //conver from int -> float64
	var f int8 = int8(e) //conver from int16 -> int8 with unexpected value (due to byte size)
	g := int(b * 4) //multiplication first, then conversion
	h := int(b) * 4 //converion first (loss of precision) then muliplication


	fmt.Println(a, b, c, d, e, f, g, h)
}