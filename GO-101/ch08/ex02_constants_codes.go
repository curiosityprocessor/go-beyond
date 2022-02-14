package main

import "fmt"

const User int = 0
const Admin int = 1

func ConstantCodes() {
	PrintUserType(User)
	PrintUserType(Admin)
}

func PrintUserType(userType int) {
	fmt.Println(userType)
}