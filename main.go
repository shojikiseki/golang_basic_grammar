package main

import (
	"fmt"
)

// 初期設定などで使うケースが多い
// func init() {
// 	fmt.Println("Init")
// }

// 【出力(fmt)】
// func bazz() {
// 	fmt.Println("Bazz")
// }

// Hello World!
// func main() {
// 関数の読み込み
// bazz()
// 文字列の連結
// fmt.Println("Hello World!", "sample string")
// }

// 【os/user, time】
// func main() {
// 	fmt.Println("Hello world!", time.Now())
// 	fmt.Println(user.Current())
// }

// 【変数代入】
func main() {
	// 値を入れなかった場合は初期値が入力される
	// var (
	// 	i    int
	// 	f64  float64
	// 	s    string
	// 	t, f bool
	// )

	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "hello!"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)

	// 変数宣言の省略形
	xi := 1
	xf64 := 1.2
	xs := "hello!"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}
