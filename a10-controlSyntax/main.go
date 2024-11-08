package main

import (
	"fmt"
	"strings"
	"time"
)

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

	// A10.2 for
	// Goのforでよく使うのはrangeを用いたスライスやMAPの全取得ループ
	// スライスはインデックスと値、MAPはキーと値がループ変数に格納される

	// rangeを利用したfor
	sketches := []string{"Dead Parrot", "Killer joke", "Mona Lisa", "Spanish Inquisition", "Spam"}
	for i, s := range sketches {
		fmt.Println(i, s)
	}

	// 変数を1つだけ書いた場合はインデックス・キーが取り出される。
	for i := range sketches {
		fmt.Println("i =", i)
	}
	// 値だけを利用したい場合、ブランク識別子_を使って読み飛ばす。
	// ブランク識別子_を使うと、未定義エラーにならない。読み飛ばす用途として明示できる。
	for _, s := range sketches {
		fmt.Println("s =", s)
	}

	// ループ内ではbreakを使った強制終了、continueを利用した処理のスキップを実装可能
	for _, s := range sketches {
		//スケッチ名がKから始まっている場合,読み飛ばす
		if strings.HasPrefix(s, "K") {
			continue
		}
		//スケッチ名がnで終わっている場合,ループ終了
		if strings.HasSuffix(s, "n") {
			break
		}
		fmt.Println(s)
	}

	// bool型を１つだけ持つforループで、条件を満たす間は回り続ける。
	counter := 0
	for counter < 10 {
		fmt.Println("counter =", counter)
		counter += 1
	}

	// すべて省略するループは無限ループになる。
	end := time.Now().Add(time.Second)
	for {
		fmt.Println("time =", time.Now())
		if end.Before(time.Now()) {
			break
		}
	}

	// ループ変数を使った、プリミティブ(原始的)なforループ
	for i := 0; i < 10; i++ {
		fmt.Println("10回繰り返す")
	}

	// A10.3 switch
	var s string = "running"
	switch s {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("その他")
	}

	// ほかの言語と異なるのが、1つのケースにマッチした後はswitchを抜ける所。
	// もし次の節に処理を継続したい場合、fallthroughを記載する
	switch s {
	case "running":
		fallthrough // runningと記載しても"run"へ続く
	case "run":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("その他")
	}

}
