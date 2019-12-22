package main

import (
	"fmt"
)

func main() {
	v := "GoLang"
	if v == "golang" {
		fmt.Println("o")
	} else if v == "GOLANG" {
		fmt.Println("Oops!")
	} else {
		fmt.Println("x")
	}
	// 変数vはif文の内外からアクセスできる
	fmt.Println(v)

	// if v := "GoLang"; v == "golang" {
	// 	fmt.Println("o")
	// } else if v == "GOLANG" {
	// 	fmt.Println("Oops!")
	// } else {
	// 	fmt.Println("x")
	// }
	// // 変数vのネームスコープがif文内だけになるので、外からアクセスするとエラーになる
	// fmt.Println(v)
}
