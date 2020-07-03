package main
import (
	"fmt"
	"math"
)
const s string = "constant"
func main() {
	// 常量 (不能第二次赋值)
	fmt.Println(s)

	const n = 500000

	var str = "hyz"
	str = "huangyuze"

	fmt.Println(str)

	const d = 3e20 / n
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}