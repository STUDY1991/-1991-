package main

/*
	数值（Number）
		整数类型
			有符号整数  int8 int16 int32（rune类型）  int64
						rune类型用来表示unicode的代码点，就是unicode字符所对应的整数
			无符号整数  uint8(byte字节类型)  uint16 uint32 uint64
						byte类型表示字节类型
			至于类型后面的数字8/16/32/64，是用来表示这个类型数据的位不同，
				位越大能表示的整数范围就越大，比如用N位来表示的整数，如果
				是有符号整数，其范围为：-2^（N-1）~2^（N-1）-1
				如果是无符号整数，其能表示的范围为：0~2^（N-1）

		浮点数类型   float16（单精度）    float32（双精度）
		虚数类型	complex64  complex128

		注：GO还定义了三个依赖系统{
				"uint  int   uintptr "
				因为值32位系统和64位系统上用来表示这些类型的位数是不一样的

			对于32位操作系统
				uint = uint32
				int = int32
				uintptr 位32位的指针
			对于64位操作系统
				uint = uint64
				 int = int64
				uintptr 为64位的指针

		}

*/
import "fmt"

const (
	iint1 int8  = 1
	iint2 int8  = 2
	iint3 int16 = 3
	iint4 int16 = 4
)
const (
	if1 float32 = 1.0
	if2 float64 = 2.0
)
const (
	icpx1 complex64  = 2
	icpx2 complex128 = 3
)

func main() {
	fmt.Println("iint1+iint2", iint1+iint2)
	fmt.Println("iint3-iint4", iint3-iint4)
	fmt.Println("iint3*iint4", iint3*iint4)
	fmt.Println("iint1/iint2", iint1/iint2)
	fmt.Println("iint4%iint3", iint4%iint3) //求余运算为整型独有，数值类型此外还包括+-*/等操作

	fmt.Println("define：字符类型（String）Go的字符串是由单个字节连接起来的，对于汉子而言，通常由多个字节组成")
	fmt.Println("传统的字符由字符组成，go的字符串不同，是由字节构成的")
	fmt.Println("字符串的表示用双引号“”或者‘’号来表示")
	fmt.Println("双引号之间的转义字符会被转义，而‘’号之间的转义字符保持原样不变")
	var a = "hello world"
	var b = 'a'
	var c = " haha~"
	fmt.Println("--打印原字符串--------")
	fmt.Println(a)
	fmt.Println("--打印字符串长度--------")
	fmt.Println(len(a))
	fmt.Println("--打印字符串第0个元素--------")
	fmt.Println(a[0])
	fmt.Println("---字符串拼接-------")
	fmt.Println(a + c)
	fmt.Println("----rune类型------")
	fmt.Println(b)
}
