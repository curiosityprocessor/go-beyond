package main

import "fmt"

type Student struct {
	Name string
	Class int
	Id int64
	Score float64
}

func Structs() {
	var student Student //declare and initialise with default values
	student.Name = "cp" //set field values
	student.Class = 2022
	student.Id = 20220101001
	student.Score = 92.75

	fmt.Println("Name:\t", student.Name)
	fmt.Println("Class:\t", student.Class)
	fmt.Println("Id:\t", student.Id)
	fmt.Println("Score:\t", student.Score)

	fmt.Println("============")

	var student2 Student = Student{"dec init", 2021, 20210709054, 84.8} //declare and initialise with given values
	fmt.Println("Name:\t", student2.Name)
	fmt.Println("Class:\t", student2.Class)
	fmt.Println("Id:\t", student2.Id)
	fmt.Println("Score:\t", student2.Score)

	fmt.Println("============")

	var student3 Student = Student{
		"multi line",
		2020,
		20201025011,
		96.7, //when initialised in multi-line, last field value must contain comma ,
	}
	fmt.Println("Name:\t", student3.Name)
	fmt.Println("Class:\t", student3.Class)
	fmt.Println("Id:\t", student3.Id)
	fmt.Println("Score:\t", student3.Score)

}