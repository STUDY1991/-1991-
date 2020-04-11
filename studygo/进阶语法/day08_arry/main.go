package main

import (
	"fmt"
)

//数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，
//这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
//var balance = [10] float64
//初始化数组
var balance = [5]float64{100.0, 2.0, 3.0, 4.0, 5.0}
var iarry = [5]int{1, 2, 3, 4, 5}

func main() {
	fmt.Println(balance)
	//遍历数组当中的元素
	for i, v := range balance {
		fmt.Printf("balance[%d] = %v\n", i, v)
	}
	//通过索引改变数组元素的值
	balance[2] = 10
	fmt.Println(balance)
	//通过索引访问数组元素，用变量接收打印输出
	var salary = balance[0]
	fmt.Println(salary)
	for k := 0; k < 5; k++ {
		fmt.Printf("iarry[%d] = %d\n", k, iarry[k])
	}
}
