package main

import "fmt"

func variableTypes() {
	/*
		[number types]
		uint: unsigned integer
		int: signed integer
		float: floating point numbers

		* Optional bit size follows number types
		ex) uint8: unsigned 1byte integer
		ex) int64: signed 8byte integer

		* Alias
		byte: alias of unint8
	*/
	var state uint = 0
	var views int = 1234567890123456789

	fmt.Println("current state:", state, "how many views?", views)

	/*
		[other types]
		bool: true/false
		string: string literal
		array
		slice
		pointer
		function
		interface
		...
	*/
	var isOpen bool = true
	var msg string = "is shop open?"

	fmt.Println(msg, isOpen)
}