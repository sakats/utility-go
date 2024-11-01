package main

import "fmt"

func main() {
	// A10.1 IF
	// if文の基礎的な書き方
	var statusCode int = 404
	if statusCode == 200 {
		fmt.Println("no error")
	} else if statusCode < 500 {
		fmt.Println("client error")
	} else {
		fmt.Println("server error")
	}

	// 変数宣言とセットの利用
	// この条件文節の中で新しく変数に設定した変数は、ブロック内でのみ利用可能
	// 上記の特性を用いて、データ取得とチェックを同時に可能

	// 数字から参照可能な名前があるとする
	cache := map[int]string{
		100: "satoshi",
		220: "sasuke",
	}
	// ※※ 難しいけど重要 ※※
	// キーの有無チェックとデータ取得
	input := 100                        // 画面からの入力を想定
	if result, ok := cache[input]; ok { // コロンの後はbool型が必要
		fmt.Println("cached value :", result) //もし変数okがtrueなら実行
	}

}
