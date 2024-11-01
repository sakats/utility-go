package main

import "fmt"

func main() {
	//a8 スライス
	// 配列やスライスを総称して複合型という。スライスのほうが使われる。

	// あまり使われない配列
	// 要素数3の整数の配列
	var nums [3]int = [3]int{1, 2, 3}
	// 要素の値を取り出して表示
	fmt.Println(nums[0])
	// 長さの取得
	fmt.Println(len(nums))

	// スライスは配列を参照する、DBでいうビューのような構造
	// 裏には配列があり、参照している配列・要素の長さ・容量の限界地が保存されている

	// スライスの変数宣言
	var nums1 []int
	// 1,2,3の要素をもつスライスを作成して代入
	nums2 := []int{1, 2, 3}

	// あるいは既存の配列やスライスからも範囲アクセスでスライス作成が可能
	nums3 := nums[0:2]  //配列numsから1~3番目の要素で作成
	nums4 := nums2[1:3] //スライスから作成

	//スライスを格納した変数の出力
	fmt.Println("nums3 =", nums3)
	fmt.Println("nums4 =", nums4)

	// 配列と同様にブラケットで要素取得可能
	fmt.Println("nums2[1] =", nums2[1])

	// 範囲外アクセスはパニックが発生
	// fmt.Println("nums2[100] =", nums2[100])
	// -> panic: runtime error: index out of range [100] with length 3

	// 要素の割り当ても可能
	nums2[1] = 100

	// 長さも取得可能
	fmt.Println("nums1 len is", len(nums1))

	// スライスに要素を追加
	// 再代入が必要
	nums2 = append(nums2, 4)

	/* ※スライスの原理
	厳密にはスライスに値を追加するという説明は間違っている。
	appendメソッドではスライスの背景にある配列の要素がなくなると、
	2倍の容量のメモリを確保してデータを乗せ換えて、新しいスライスを返す。
	Goの言語使用的には可変長配列は存在しないが、スライスのこの挙動で実現している。
	*/

	// A9 MAP
	// いわゆるmap,辞書,ハッシュと呼ばれる構造。マップも複合型に含まれる。
	// 型は map[キーの型]mapの型 という形で表現される
	// mapの利用には初期化が必要であり、ブレース{}をつけるか、make関数で作成する。
	// 値の割り当てと取得にはブラケット[]を利用する。

	// 数字がキーで値が文字列のマップに、HTTPステータスを格納
	// ブレースで作成する場合
	hs := map[int]string{
		200: "OK",
		404: "Not Found",
	}

	// makeで作成する場合
	authors := make(map[string][]string)

	// ブラケットで要素アクセス（代入）
	authors["Go"] = []string{"Satoshi Akatsuka", "Ken Ito", "Ryosuke Kameya"}

	// データ取得
	status := hs[200]
	fmt.Println(status) // OK

	// 存在しない要素にアクセスするとゼロ値が得られる
	fmt.Println(hs[0]) //panicは発生しない？

	// 存在するかどうかの情報も一緒に取得できる
	status, ok := hs[304]
	fmt.Println("status =", status) // ""
	fmt.Println("ok =", ok)         // false

}
