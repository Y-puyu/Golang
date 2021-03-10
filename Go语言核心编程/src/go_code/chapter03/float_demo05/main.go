package main

import (
	"fmt"
)

// 整数类型的使用
func main() {
	// 单精度float32
	// 双精度double32
	// 浮点数=符号位+指数位+尾数位，浮点数都是有符号的
	var price float32 = 12.11
	fmt.Println(price)

	// 浮点数精度丢失
	var a float32 = 12.00001212 // 12.000012
	var b float64 = 12.00001212 // 12.00001212
	fmt.Println(a, b)
	
	// 浮点类型默认为float64
	var c = 12.12121
	fmt.Printf("%T\n", c) // float64

	// 表示浮点数的两种方式
	d := 1.234
	e := .567
	fmt.Println(d, "\n", e)

	
	// 科学计数法
	g := 5.1234e2
	// h := 5.1234e2	// vscode插件默认将大写E改为小写e
	k := 5.1234e-2 // 0.051234
	fmt.Println(g, "\n", k)
	
}
