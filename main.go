package main

import "fmt"

var iint int = 1                      //定义一个int类型的变量iInt  赋值为1
var istring string = "helloworld ~~~" //定义了一个字符串变量istring  值是“hello world~~~”
const hw string = "hello ,world !!!"  //定义了一个常量hw字符串

func main() {
	fmt.Printf("hello world @_@\n")
	fmt.Printf("iint:%d", iint)
	fmt.Println("")
	fmt.Println(hw)
	fmt.Println(iint)
	fmt.Printf(istring)
	fmt.Print(hw)
}

/*
println会根据你输入格式原样输出，自动换行，
printf需要格式化输出并带输出格式；
*/
