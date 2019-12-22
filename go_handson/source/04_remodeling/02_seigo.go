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

	ccnt, wcnt := 0, 0
	for _, v := range data {
		fmt.Printf("%s\n", v)

		var ans string;
		fmt.Print("input: ")
		fmt.Scanln(&ans)

		if ans == v {
			fmt.Println("o")
			ccnt++
		} else {
			fmt.Println("x")
			wcnt++
		}

		fmt.Println()
	}
	fmt.Printf("[正誤] 正解：%d, 誤り：%d\n", ccnt, wcnt)
}
