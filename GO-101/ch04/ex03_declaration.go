package main

import "fmt"

func variableDeclaration() {
	/*
		variable declaration
	*/
	var a int = 3 // basic
	var b, c int =  2, 1 //multiple declaration
	var d int //no explicit initial value. initialised with default value for current data type
	var e = true //no explicit type. type is inferred from the assigned value
	f := "GO" //no explicit var keyword, type. used with shorthand sytax :=
	fmt.Println(a, b, c, d, e, f)
}