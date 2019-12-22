package main

import (
	"fmt"
)

var (
	example = []string{
		"golang",
		"handson",
		"in",
		"umeda",
	}
)

func main()  {
	var data []string
	data = example

	for i, v := range data {
		fmt.Printf("Q%d: %s\n", i + 1, v)

		var ans string;
		fmt.Print("input: ")
		fmt.Scanln(&ans)

		if ans == v {
			fmt.Println("o")
		} else {
			fmt.Println("x")
		}

		fmt.Println()
	}
}
