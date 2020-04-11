package main

import "fmt"

//方式1，一般变量的定义方式   var  变量名 类型 = 变量值
var a string = "hello world!" 
//a为全局变量 相反在函数内部定义的变量称为局部变量，只能在函数内部访问使用

func main() {
	fmt.Printf("a:%s,%T", a, a)
	var b string //方法2  先定义   后赋值
	b = "\nlet's go"
	fmt.Println(b)
	c := "Great!"  //方法3  变量名:= 变量值  （快捷方式）  只能在函数内部使用这种方式
	fmt.Println(c) //变量c 的值是由Go语言自己推断他的类型  适用于知道变量的初始值
}
