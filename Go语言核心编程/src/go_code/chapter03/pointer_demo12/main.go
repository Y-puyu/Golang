package main

import (
	"fmt"
)

// 指针类型
func main() {
	var i int = 10
	// i 的地址是 &i
	fmt.Println(&i) // 0xc0000a2058

	var p *int = &i
	fmt.Println(*p)

	fmt.Printf("%v\n", p)  // 0xc0000a2058
	fmt.Printf("%v\n", &p) // 0xc0000cc020

	var num int = 9
	fmt.Println(num, &num) // 9 0xc0000a2098

	var n *int = &num
	*n = 10
	fmt.Println(n, *n, num) // 0xc0000a2098 10 10

	/* 练习
	var a int = 300
	var ptr *int = a	// 不能把值赋给一个指针变量

	var a int = 300
	var ptr *float = &a	// 要求类型相同
	*/

	// 值类型，基本的数据类型 int, float, bool, string, 数组，结构体。都有对应的指针类型，
	// 引用类型，指针，slice 切片，map，管道 chan，interface 等都是引用类型
}
