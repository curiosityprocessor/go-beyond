package main

import "fmt"

type Student struct {
	Name string
	Class int
	Id int64
	Score float64
}

func main() {
	var student Student
	student.Name = "cp"
	student.Class = 2022
	student.Id = 20220101001
	student.Score = 92.75

	fmt.Println("Name:\t", student.Name)
	fmt.Println("Class:\t", student.Class)
	fmt.Println("Id:\t", student.Id)
	fmt.Println("Score:\t", student.Score)

}