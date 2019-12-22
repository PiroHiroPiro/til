package main

import (
	"bufio"
	"fmt"
	"os"
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
		texts = append(texts, text)
	}

	if err := scan.Err(); err != nil {
		return nil, err
	}

	return texts, nil
}

func main() {
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

	// 指定されたファイル名の中身を出力
	if texts, err := output(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
	} else {
		for _, text := range texts {
			fmt.Println(text)
		}
	}
}
