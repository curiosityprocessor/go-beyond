package main

import "fmt"

func main() {
	if grade, isPass := getGrade(); isPass && grade == 100 { //if with init statement
		fmt.Println("You aced it!")
	} else if isPass && grade > 90 { //accessing variables declared in if init
		fmt.Println("You did well")
	} else if isPass {
		fmt.Println("You passed")
	} else {
		fmt.Println("Better luck next time!")
	}

	// fmt.Println("Your grade is....", grade) //variable out of scope - error: undefined: grade
}

func getGrade() (int, bool) {
	return 97, true
}