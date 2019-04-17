package main

// 【基本的な関数】
// ---------------------------------------------------------------
// 基本ケース
// func hoge() {
// 	処理
// 	処理
// }

// 初期設定などで使うケースが多い
// func init() {
// 	fmt.Println("Init")
// }

// 【出力(fmt)】
// ---------------------------------------------------------------
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
// ---------------------------------------------------------------
// func main() {
// 	fmt.Println("Hello world!", time.Now())
// 	fmt.Println(user.Current())
// }

// 【変数宣言】
// ---------------------------------------------------------------
// varの形式であれば関数外でも使用可
// var (
// 	i    int     = 1
// 	f64  float64 = 1.2
// 	s    string  = "hello!"
// 	t, f bool    = true, false
// )

// 変数宣言の省略形
// ※省略する場合は関数内でのみ使用可
// func foo() {
// 	xi := 1
// 	var xf32 float32 = 1.2
// 	xs := "hello!"
// 	xt, xf := true, false
// 	fmt.Println(xi, xf32, xs, xt, xf)
// 	fmt.Printf("%T\n", xf32)
// 	fmt.Printf("%T\n", xi)
// }

// 【変数代入】
// func main() {
// 値を入れなかった場合は初期値が入力される
// var (
// 	i    int
// 	f64  float64
// 	s    string
// 	t, f bool
// )

// 	fmt.Println(i, f64, s, t, f)
// 	foo()
// }

// 【const（不変変数）】
// ---------------------------------------------------------------
// const Pi = 3.14
// const (
// 	Username = "test user"
// 	Password = "test pass"
// )

// overflowした場合は出力時に -1 してもダメ
// var big int = 9223372036854775807 + 1
// ただしconstの場合はoverflowでも大丈夫
// const big = 9223372036854775807 + 1

// func main() {
// 	fmt.Println(Pi, Username, Password)
// 	fmt.Println(big - 1)
// }

// 【数値型】
// ---------------------------------------------------------------
