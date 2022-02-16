package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	var str string = "Hello Go World"
	var slice []byte = []byte(str)
	
	//use unsafe and reflect package to transform string into construct
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceheader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringheader.Data)
	fmt.Printf("slice:\t%x\n", sliceheader.Data)
}