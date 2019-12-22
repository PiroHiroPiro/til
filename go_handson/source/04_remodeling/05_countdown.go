package main

import (
	"fmt"
	"time"
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
	ln := len(data)

	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	gstm := time.Now()
	ccnt, wcnt := 0, 0
	for i, v := range data {
		fmt.Printf("[%d/%d] %s\n", i + 1, ln, v)

		var ans string;
		fmt.Print("input: ")
		stm := time.Now()
		fmt.Scanln(&ans)
		ptm := time.Since(stm)

		if ans == v {
			fmt.Println("o")
			ccnt++
		} else {
			fmt.Println("x")
			wcnt++
		}
		fmt.Printf("%.3f(sec)\n", ptm.Seconds())
		fmt.Println()
	}
	fmt.Printf("[正誤] 正解：%d, 誤り：%d\n", ccnt, wcnt)
	gptm := time.Since(gstm)
	fmt.Printf("Total time: %.3f(sec)\n", gptm.Seconds())
}
