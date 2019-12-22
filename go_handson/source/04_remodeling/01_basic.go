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
	data := example

	for _, v := range data {
		fmt.Printf("%s\n", v)

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
