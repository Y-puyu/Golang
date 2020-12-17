package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型和string的转换
func main() {
	var n1 int = 99          // %d  查官方文档即可
	var n2 float32 = 23.1123 // %t 布尔值， %b二进制，%v默认格式
	var n3 bool = true
	var n4 byte = 'm'

	var str string // 空的str

	// 方式一：使用fmt.Sprintf方法(这个方法用的十分多)
	str = fmt.Sprintf("%d", n1)
	fmt.Printf("%T %q\n", str, str) // %v也行，%q将字符串加双引号输出（原始定义不是这样）

	str = fmt.Sprintf("%f", n2)
	fmt.Printf("%T %q\n", str, str)

	str = fmt.Sprintf("%t", n3) // %t 布尔类型
	fmt.Printf("%T %q\n", str, str)

	str = fmt.Sprintf("%c", n4)
	fmt.Printf("%T %q\n", str, str)

	fmt.Printf("%q\n", n1) // 输出 'c'。将单引号括起来的'99'输出为字符'c'

	// 方式二：使用  strcnov 函数
	var n5 int = 99
	var n6 float64 = 23.1212
	var n7 bool = true

	str = strconv.FormatInt(int64(n5), 10) // 函数第一个参数需要接收int64
	fmt.Printf("%T %q\n", str, str)

	// 说明：'f' 即常规的小数点格式输出，10 指小数保留多少位，64 表示这个小数是float64的
	str = strconv.FormatFloat(n6, 'f', 10, 64)
	fmt.Printf("%T %q\n", str, str)

	str = strconv.FormatBool(n7)
	fmt.Printf("%T %q\n", str, str)
}
