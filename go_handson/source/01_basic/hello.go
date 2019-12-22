// 実行させるならmainを指定
package main

// 外部パッケージを指定
import (
	"fmt"
)

// 同一パッケージ内で使用できる変数を定義
var (
	// 変数msg1は宣言+初期化（値から推測してstring型）
	msg1 = "Hello"
	// 変数msg2は明示的にstring型で宣言
	msg2 string
)

// main()関数：エントリーポイント
func main()  {
	msg2 = "world"
	// Println()関数は、引数が複数だと半角空白文字で区切って標準出力に表示
	// fmtパッケージのPrintln()関数は名前が大文字で始まっているので呼び出せる
	fmt.Println(msg1, msg2)
}
