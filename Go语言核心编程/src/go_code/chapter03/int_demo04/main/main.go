package main

import (
	"fmt"
	"go_code/chapter03/int_demo04/mo"
	"unsafe"
)

// 整数类型的使用
func main() {
	var i int = 1
	fmt.Println(i)

	// 测试int8的范围
	var j int = -128
	// var j int8 = -129

	// var j int8 = 128
	// var j int8 = 127
	fmt.Println(j)

	// 测试uint8的范围
	var k uint8 = 0
	// var k uint8 = 255
	// var k uint8 = -1
	// var k uint8 = 256
	fmt.Println(k)

	// int, uint, rune, byte
	// int,uint根据64 32位版本走
	// rune相当于int32 	-2^31~2^31-1
	// byte相当于uint8 	0~255。在存单个字母、字符的时候采用byte存储
	var a int = 2147483647
	fmt.Println(a)

	// 整形默认类型为int类型
	var n1 = 100 // n1是什么类型
	// 用一个方法，可以查看某个变量的数据类型
	// fmt.Printf() 用于做格式化输出
	fmt.Printf("%T\n", n1) // 将打印int，默认为int，至于是int32还是int64根据电脑有关

	// 怎么查看变量占用字节大小？
	var n2 int64 = 10
	// unsafe.Sizeof(n1) 是unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("%T %d %T %d\n", n1, unsafe.Sizeof(n1), n2, unsafe.Sizeof(n2))

	// 引包
	fmt.Println(mo.HearName)
}
