package main

import "fmt"

//方式1.定义常量  使用const关键字
const name string = "李正洪"
const department string = "iMB"

//方式2.快捷定义常量
const (
	tel = 18383123662
	iD  = "C4853062"
)

func main() {
	fmt.Printf("name:%s\ndepartment:%s\n", name, department)
	fmt.Printf("tel:%d\nId:%s\n", tel, iD)
	zc()
}

// 定义一个计算周长的函数
func zc() {
	var (
		a = 10
		b = 5
	)
	var zc = (a + b) * 2
	fmt.Printf("周长= (%d+%d)*2=%d", a, b, zc)
}
