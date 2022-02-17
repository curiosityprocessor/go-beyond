package main

import (
	"fmt"
	"bufio"
	"os"
	"guess_my_number/generator"
)

var stdin = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("======= START ======")
	
	num := generator.RandomNoGenerator()
	fmt.Println(">>> DEBUG num:", num)

	fmt.Println("Guess my number")
	input, err := InputValue()
	if err != nil {
		fmt.Println("Something went wrong... Try again!")
	} else {
		fmt.Println(">>> DEBUG input:", input)
	}
	
	fmt.Println(num == input)
	
	fmt.Println("======== END =======")
}

func InputValue() (int, error) {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		stdin.ReadString('\n')
	}
	return input, err
}