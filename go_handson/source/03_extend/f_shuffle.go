package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
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

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func shuffle(src []string) []string {
	idxs := rand.Perm(len(src))

	// 更新
	// dst := make([]string, len(src))
	// for i, v := range idxs {
    //     dst[i] = src[v]
	// }

	// 追加
	dst := make([]string, 0)
	for _, v := range idxs {
		dst = append(dst, src[v])
	}

	return dst
}

func randomCase(src string) string {
	dst := ""
	for _, c := range src {
		s := string(c)
		idx := rand.Intn(2)
		if idx == 0 {
			dst = dst + strings.ToUpper(s)
		} else {
			dst = dst + strings.ToLower(s)
		}
	}

	return dst
}

func main()  {
	var data []string

	// 引数の数をチェック
	if len(os.Args) == 1 {
		data = example
	} else if len(os.Args) == 2 {
		filename := os.Args[1]
		// ファイルの存在をチェック
		if _, err := os.Stat(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// 指定されたファイル名の中身を出力
		if texts, err := output(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			data = texts
		}
	} else {
		fmt.Fprintln(os.Stderr, "Usage: ./e_args [filename]")
		os.Exit(0)
	}

	data = shuffle(data)
	ln := len(data)

	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	ccnt, wcnt := 0, 0
	ctime := 0.0
	for _, v := range data {
		w := randomCase(v)
		fmt.Printf("%s\n", w)

		var ans string;
		fmt.Print("input: ")
		stm := time.Now()
		fmt.Scanln(&ans)
		ptm := time.Since(stm)
		ctime = ctime + float64(ptm.Seconds())

		if ans == w {
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
