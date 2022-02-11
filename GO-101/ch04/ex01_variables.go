package main

import "fmt"

func variables() {
	/*
		var: keyword to declare variable
		a: variable name
		int: variable type
		10: initial value
	*/
	var a int = 10
	var msg string = "String variable"
	fmt.Println(msg, a)

	a = 20
	msg = "New String"
	fmt.Println(msg, a)
}