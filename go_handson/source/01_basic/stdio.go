package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("入力は? > ")
	var ans string

	fmt.Scanln(&ans)
	// 変数名に&がついていないと値が書き換わらない
	// https://github.com/golang/go/blob/4d5bb9c60905b162da8b767a8a133f6b4edcaa65/src/fmt/scan.go#L953
	// fmt.Scanln(ans)

	fmt.Printf("入力は %s です\n", ans)

	// 改行させる場合は、改行コードを含める
	fmt.Print("golang ")
	fmt.Print("2019\n")

	// 最後にf付きは、フォーマット形式
	fmt.Printf("%s %d\n", "golang", 2019)

	// 最後にln付きは、文末に改行コードが付く
	fmt.Println("golang", "2019")

	// 最後にF付きは、第一引数に出力先を指定（ここでは標準エラー）
	fmt.Fprintln(os.Stderr, "golang", "2019")

	// 最後にS付きは、文字列を返す
	s := fmt.Sprintf("%s %d", "golang", 2019)
	fmt.Println(s)
}
