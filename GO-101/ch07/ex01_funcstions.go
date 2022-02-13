package main

import "fmt"

func main() {
	c:= Add(2, 5)
	fmt.Println(c)
}

/*
	func: keyword for declaring functions
	Add: function name
	(a int, b int): parameters
	) int {: return type
	{}: function code block

	** GO supports pass by value **
	parameters are copied values not addresse references
*/
func Add(a int, b int) int {
	return a + b
}