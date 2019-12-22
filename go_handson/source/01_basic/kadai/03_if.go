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

	for _, v := range data {
		fmt.Print(v, "\t")
		if v == "in" {
			fmt.Println("o")
		} else {
			fmt.Println("x")
		}
	}
}
