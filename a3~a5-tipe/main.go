package main

import (
	"fmt"
	"strconv"
)

func main() {
	// A3 命名
	// 変数名は基本ローワーキャメル
	var firstName = "satoshi"
	fmt.Println(firstName)

	// アッパーキャメルで書くとグローバル要素になる
	var LastName = "Akatsuka"
	fmt.Println(LastName)

	// A4 コメント
	// 行コメント
	/*
		ブロックコメント
	*/

	// A5 型と変換
	// Goではリテラルの時点で型を持たない。変数に代入した際に型が決まる。
	// キャストでは暗黙の型変換がなく、明示的に型変換の必要がある。
	var i int = 123

	// 数値同士の変換は、カッコでくくり型を前置する。
	var f float64 = float64(i)

	// 64bitOSで、64bitのintと、int64も明示的な変換が必要。
	// (元からint型で生成した変数の場合、PCが64bitでも、64bit型とは扱いが異なるという意味)
	var i64 int64 = int64(f)
	fmt.Println(i64)

	// boolへの変換は比較演算子を使う
	var b bool = i != 0
	fmt.Println(b)

	// 数値→文字列の変換
	var s string
	s = strconv.FormatInt(i64, 10) // 第二引数は何進数に変換するか（10 → 10進数）
	fmt.Println("s =", s)

	// strconvの数値入力はint64, uint64, float64のみ、それ以外はキャストが必要
	s = strconv.FormatInt(int64(i), 10)
	fmt.Println("s =", s)

	// 文字列→数値の変換
	// Parse系は変換失敗時にエラーを返す
	fl, err := strconv.ParseFloat("12345.64", 64) // 第二引数はビット精度（64 → float64）
	fmt.Println(fl, err)

}

// GoDocについて
// 改行なしで関数や型の宣言の直前にコメントをつけると、
// それらの要素の説明に利用される（GoDocから参照される）
func CommentedFunc() {
	fmt.Println("anything")
}
