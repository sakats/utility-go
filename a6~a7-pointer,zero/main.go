package main

import (
	"fmt"
)

func main() {
	//a6 pointer

	// リテラルを変数に代入すると、PC上のメモリに格納されてアドレスが生まれる
	var i int = 10
	// 変数のポインタ型はアスタリスクを前置する, デフォルトはnil
	var p *int
	// iのポインタをpに格納する
	p = &i
	// p経由で,iのアドレスと値を見る
	fmt.Println("p =", p)   // p = 0xc00008c098
	fmt.Println("*p =", *p) // *p = 10

	//a7 zero
	// 変数宣言のみを行った変数は「ゼロ値」で初期化される。
	// 数値は0,boolならfalse,文字列なら"",ポインターやインターフェースならnil

	var z int
	fmt.Println("z =", z) // z = 0

	var b bool
	fmt.Println("b =", b) // b = false

	var s string
	fmt.Println("s =", s) // s =

	var po *int
	fmt.Println("po =", po) // po = <nil>
}
