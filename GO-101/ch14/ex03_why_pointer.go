package main

import "fmt"

type Data struct {
	value int
	data [200] int
}

func main() {
	var data Data //data.value = 0, data.data[100] = 0

	ChangeData(data) //changes in function call are not applied to original variable

	fmt.Println("===== pass by value =====")
	fmt.Println("value: ", data.value) //still 0
	fmt.Println("data[100]: ", data.data[100]) //still 0

	ChangeDataByPointer(&data) //memory address as parameters

	fmt.Println("===== pass by reference =====")
	fmt.Println("value: ", data.value) //now 999
	fmt.Println("data[100]: ", data.data[100]) //now 100

}

//parameter values are **copied**
//pass by value, NOT pass by reference
func ChangeData(data Data){
	data.value = 999
	data.data[100] = 100
}

//pointers passed as parameters
//TECHNICALLY pass by value of pointer
//works like pass by reference
func ChangeDataByPointer(data *Data) {
	data.value = 999
	data.data[100] = 100
}