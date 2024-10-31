package main

import (
	"fmt"
)

func main() {
	//A2 literal
	//1     //定数
	//1.23  //浮動小数点
	//"abc" //文字列
	//`今日も
	//がんばろう` //文字列（バッククォートで改行可）
	//true //bool
	//nil  //無効な参照先を表す値

	// 整数型の宣言
	var num1 int = 123
	// 右辺で型が決まるなら省略可
	var num2 = 123

	// 変数は利用しないとコンパイルエラー
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)

	// シャドーイング（変数の二重宣言）
	x := 10
	condition := true
	// if分でネストするとシャドーイング可
	if condition {
		// コードブロック内で変数xを再定義
		x := 20
		fmt.Println("x =", x)
	}
	// シャドーイングの影響を受けないため10が出力される
	fmt.Println("x =", x)
}
