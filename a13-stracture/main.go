package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// A13 構造体
	// いくつかのデータをまとめて使えるようにしたのが構造体
	// ブレース{}の中にフィールド（メンバーとなる変数）を列挙する
	// 構造体のフィールドも大文字にしないと外部パッケージから利用できない
	// 他言語で言うクラスとは異なり、中にメソッドを定義はできない

	// 構造体の定義例
	type Book struct {
		Title      string
		Author     string
		Publisher  string
		ReleasedAt string
		ISBN       string
	}

	// 構造体の利用
	// インスタンス生成(フィールドはすべてゼロ値で初期化される)
	var b Book
	fmt.Println("b :", b)

	// フィールドを初期化しながらインスタンス生成
	b2 := Book{
		Title: "三日間の幸福",
	}
	fmt.Println("b2 :", b2)

	// フィールドを初期化しながらインスタンス生成
	// 変数にはポインターを格納
	b3 := &Book{
		Title: "Die with Zero",
	}
	fmt.Println("b3 :", b3)

	// 下記構造体の活用に続く
	roadJson()

}

// 構造体の活用

// jsonタグを定義しておくと、構造体のフィールドをJSONに書き出したり、
// 逆にJSONの情報をフィールドにマッピングできる
// jsonを扱うには"encoding/json"パッケージをインポートする

// book.jsonの中身を構造体にロードする
type Book2 struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	ReleasedAt string `json:"resease_at"`
	ISBN       string `json:"isbn"`
}

// JSONをロードする関数
func roadJson() {
	f, err := os.Open("book.json")
	if err != nil {
		log.Fatal("file open error :", err)
	}
	d := json.NewDecoder(f)
	var b Book2
	d.Decode(&b)
	fmt.Println("b :", b)
}
