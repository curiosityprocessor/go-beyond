package main

import "fmt"

func Strings() {
	str1 := "Hello\t'GO'\tworld!\n" //enables escape charaters
	str2 := `Hello 'GO' world!\n` //string literal as is

	fmt.Print(str1)
	fmt.Print(str2)
}