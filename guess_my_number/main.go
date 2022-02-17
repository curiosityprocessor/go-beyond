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
	
	actual := generator.GenerateRandomNum()
	fmt.Println(">>> DEBUG num:", actual)

	fmt.Println("Guess my number")
	for {
		guess, err := ReceiveInt()
		if err != nil {
			fmt.Println("Not a valid number... Try again!")
			continue
		}
		
		isCorrect, message := isGuessCorrect(actual, guess)
		fmt.Println(message)
		
		if(isCorrect) {
			break
		}
	}
	
	fmt.Println("======== END =======")
}

func isGuessCorrect(actual, guess int) (bool, string) {
	switch {
		case guess == actual:
			return true, "Correct!"
		case guess > actual:
			return false, "Down"
		case guess < actual:
			return false, "Up"
		default:
			return false, "Try again"
	}
}

func ReceiveInt() (int, error) {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		stdin.ReadString('\n')
	}
	return input, err
}