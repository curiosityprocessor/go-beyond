package main

import "fmt"

var globalVar string = "global scope" //scope: global

// func main() {
// 	variableScope()
// 	fmt.Println("accessible variables 3:", globalVar)
// }

func variableScope() {
	var functionScopeVar string = "function scope" //scope: func variableScope(){}
	{
		var scopeVar string = "code block" //scope: code block {}
		fmt.Println("accessible variables 1:", scopeVar, functionScopeVar, globalVar)
	}
	fmt.Println("accessible variables 2:", functionScopeVar, globalVar)

}