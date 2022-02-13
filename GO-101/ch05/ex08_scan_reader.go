package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b) //first round of input

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') //read until \n == emptying the input stream
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b) //second round of input - with emptied input stream

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') //read until \n == emptying the input stream
	} else {
		fmt.Println(n, a, b)
	}

	
}