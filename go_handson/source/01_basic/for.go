package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 4; i++{
		fmt.Print(i)
	}
	fmt.Println()

	// [:=]記述：変数iは宣言+初期化（値から推測してint型）
	// i := 1
	// for {
	// 	fmt.Print(i)
	// 	i++
	// }

	// 変数ABCをstring型のスライスで宣言して、値を設定
	var abc []string
	abc = []string{"a", "b", "c"}

	// 第2戻り値変数で添字、第2戻り値変数で値が取得できる
	for i, v := range abc {
		fmt.Println(i, v)
	}

	for _, v := range abc {
		fmt.Println(v)
	}

	// 添字だけを取得する場合は、第1戻り値変数のみ指定する（i, _でも良い）
	for i := range abc {
		fmt.Println(i)
	}
}
