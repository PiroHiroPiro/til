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

	ccnt, wcnt := 0, 0
	ctime := 0.0
	for _, v := range data {
		fmt.Printf("%s\n", v)

		var ans string;
		fmt.Print("input: ")
		stm := time.Now()
		fmt.Scanln(&ans)
		ptm := time.Since(stm)
		ctime = ctime + float64(ptm) / float64(time.Microsecond)

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
	fmt.Printf("平均入力時間：%.3f\n", ctime / float64(ln))
}
