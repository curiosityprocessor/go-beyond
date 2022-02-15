package main

import (
	"fmt"
	"unsafe"
)

/*
	Sample struct consists of total 19 bytes
*/
type Sample struct {
	A int8 //1 byte
	B int64 //8 byte
	C int8 //1 byte
	D int64 //8 byte
	E int8 //1 byte
}

/*
	Sample2 struct consist of same total 19 bytes but in efficient order
	fields with size smaller than 8 bytes are ordered consecutively to fill one 8-byte
*/
type Sample2 struct {
	A int8 //1 byte
	C int8 //1 byte
	E int8 //1 byte
	B int64 //8 byte
	D int64 //8 byte
}

func main() {
	sample := Sample{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(sample))
	// actual memory size is 40 bytes due to memory padding from memory alignment

	sample2 := Sample2{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(sample2))
	// actual memory size is 24 bytes after memeory alignment
}