package main

import ( 
	"fmt"
	"time"
)

func main() {
	fmt.Println("default for loop")
	var count int = 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	fmt.Println("================")

	fmt.Println("for loop without init statement")
	var j int = 0
	for ; j < count; j++ {
		fmt.Println(j)
	}

	fmt.Println("================")

	fmt.Println("for loop without post statement")
	for k := 0; k < count; {
		fmt.Println(k)
		k++
	}
	
	fmt.Println("================")

	fmt.Println("for loop with only conditional == while")
	var p = 0
	for p < 5 {
		fmt.Println(p)
		p++
	}

	fmt.Println("for loop with no init, conditional, post statement == infinite loop")
	var q = 0
	for {
		time.Sleep(time.Second)
		println(q)
		q++
		if q >= 5 {
			break;
		}
	}
	
	
}