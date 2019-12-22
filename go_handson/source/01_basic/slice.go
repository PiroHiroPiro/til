package main

import (
	"fmt"
)

func main() {
	// 変数exampleをstring型スライスで宣言
	var example []string
	fmt.Println("初期値は", example)

	// 値をまとめて代入（}の前に要素がなければ最後の要素にカンマが必要！）
	example = []string{
		"Golang",
		"handson",
	}
	// example = []string{"Golang", "hands-on"}
	fmt.Println("2番目の値は", example[1])

	// 値を追加するにはappend関数を使用
	example = append(example, "in", "kagawa")
	fmt.Println("追加結果は", example)

	// 連続した値
	fmt.Println(example[1:3])
	// 指定した位置以降（以前だと[:2）
	fmt.Println(example[:2], example[2:])

	// 4番目の値を変更
	example[3] = "umeda"
	fmt.Println(example)
}
