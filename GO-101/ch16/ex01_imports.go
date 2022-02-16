package main

import (
	"fmt"
	"math/rand" //packages with path. used with last path name

	"text/template" //duplicate package name( = last path) 
	htemplate "html/template" //use alias to distinguish

	_ "database/sql" //import not explicitly used but need can have a leading _ to avoid error
	//otherwise all imports must be used

	"GO-101/ch16/custom" //use custom package. package path follows module name from go mod init
)

func main() {
	fmt.Println(rand.Int())
	template.IsTrue(0)
	htemplate.IsTrue(true)
	custom.PrintCustom()
}