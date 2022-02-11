package main

import "fmt"

func main() {
	variableSize()
}

func variableSize() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b //precision lost! expected: 4,266,663.334329 computed: 4266663 error: -0.334329
	var d float32 = c * 3 //precision lost! expected: 12,799,990.002987 computed: 12,799,989 error: -1.002987

	var e float64 = float64(a) * float64(b) //computed: 4,266,663.216691017 error: 0.117637983
	var f float64 = e * 3 //computed: 12,799,989.650073051 error: 0.352913949

	fmt.Println(a, b, c, d, e, f)
}