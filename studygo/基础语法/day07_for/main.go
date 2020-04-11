package main

import "fmt"

var i int

func main() {
	for i = 0; i <= 20; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")
	jiujiu1()
	fmt.Println()
	jiafa()
}
func jiujiu1() {
	var i1 int = 1
	var i2 int = 1
	for i1 = 1; i1 < 10; i1++ {
		for i2 = 1; i2 <= i1; i2++ {
			sum := i1 * i2
			fmt.Printf("%d*%d=%d ", i2, i1, sum)
		}
		fmt.Println("")
	}
}
func jiafa() {
	// 遍历, 决定处理第几行
	for y := 1; y <= 9; y++ {
		// 遍历, 决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d+%d=%d ", x, y, x+y)
		}
		// 手动生成回车
		fmt.Println()
	}
}
