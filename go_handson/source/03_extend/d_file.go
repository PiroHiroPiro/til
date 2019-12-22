package main

import (
	"bufio"
	"fmt"
	"os"
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

func output(filename string) ([]string, error) {
	// ファイルをオープン
	f, err := os.Open(filename)
	// ファイルをオープンする時に発生したエラーハンドリング
	if err != nil {
		return nil, err
	}

	// ファイルをクローズする
	// defer文で指定された関数は自関数の終了時に実行される
	// exitされるとdefer文で指定した関数は実行されない！
	defer f.Close()

	var texts []string
	// ファイルの中身をスキャンして１行ずつ出力
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		text := scan.Text()
		// 空行は無視
		if text == "" {
			continue
		}
		texts = append(texts, text)
	}

	if err := scan.Err(); err != nil {
		return nil, err
	}

	return texts, nil
}

func main()  {
	// 引数の数をチェック
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "invalid args")
		os.Exit(1)
	}
    filename := os.Args[1]

	// ファイルの存在をチェック
	if _, err := os.Stat(filename); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

	var data []string

	// 指定されたファイル名の中身を出力
	if texts, err := output(filename); err != nil {
		data = example
	} else {
		data = texts
	}

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
		ctime = ctime + float64(ptm.Seconds())

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
