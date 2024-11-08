package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// a11 method
// main関数はgo言語のランタイムから呼ばれる関数である。
// 関数も配列･スライス･MAPと同様に複合型である。
func main() {
	// 関数の呼び出し例
	fmt.Println(calc(29, 19))
	fmt.Println(calcAge(1998, 5, 15))
	fmt.Println(calcAge(2024, 11, 9))
	test()
	fileOpen()
}

// 関数定義の例
// 関数の戻り値・引数が複数あり、同じ型が連続する場合は省略可能。後ろの型が前にもある判定。
func calc(x, y int) int {
	return x + y
}

// 戻り値に名前を付けられる。戻り値用の変数が定義済となり、returnと書くだけで戻り値を返して終了できる。
// 戻り値が複数個になったり、名前を付ける場合はかっこでくくる。
func calcAge(y int, m time.Month, d int) (age int, err error) {
	b := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	n := time.Now()
	if b.After(n) {
		err = errors.New("誕生日が未来です。")
		return
	}
	for {
		b = time.Date(y+age+1, m, d, 0, 0, 0, 0, time.Local)
		if b.After(n) {
			return
		}
		age++
	}
}

// 無名関数 test関数内で作成された無名関数mをdoCalcに渡して実行
func test() {
	// 無名関数
	// 何かのタイミングでコールバックされる処理はコールバックを受け取る場所で一緒に定義可能
	// 名前付き関数はパッケージレベルでしか作成できないが、無名関数は関数内で使える
	// 無名関数はコールバック関数として他の関数の引数や、変数に入れることができる。
	m := func(x, y int) int {
		return x * y
	}

	// 名前付きで定義した関数以外に、無名関数として定義できたものを渡せる
	doCalc(10, 20, m)
	// その場で定義した関数も渡せる
	doCalc(10, 10, func(x, y int) int {
		return x * y
	})
	// def-func5
}

// 関数を受け取る関数
func doCalc(x, y int, f func(int, int) int) {
	fmt.Println(f(x, y))
}

// a11.1 deferで後処理の関数の実行予約
// deferを使うと、その関数のブロックを抜けるタイミングで処理を実行
// pythonのwithのようなイメージ

func fileOpen() {
	f, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	// 関数のスコープを抜けたら自動でファイルをクローズ
	defer f.Close()
	io.WriteString(f, "hello World") // deferで後処理を書いてから、メインの処理を実装
}
