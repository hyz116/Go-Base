package main
import "fmt"
func main() {
   // 根据变量的初始值 推断变量类型

	// 定义一个变量
	var a = "initail"
	fmt.Println(a)

	// 定义多个变量
	var b, c int = 1,2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// var f string = "apple"
	f := "apple"
	fmt.Println(f)

	var g string
	fmt.Println(g)
}