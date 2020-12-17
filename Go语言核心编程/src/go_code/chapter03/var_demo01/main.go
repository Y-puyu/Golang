package main

import "fmt"

// 定义全局变量
var b1 = 100
var b2 = 200
var b3 = "sun"

// 全局变量声明
// 上面的声明方式，更改为一次性声明
var (
	n4 = 101
	n5 = 201
	n6 = "bool"
)

// 这样的全局变量声明是错误的
// n7, n8, n9 := 100,101,102
// var n7, n8, n9 := 100,101,102

func main() {
	// 变量定义及初始化
	//
	// 方式一：定义int变量i，赋值为10，不赋值默认为0
	var i, j int
	i = 10
	fmt.Println(i, j)

	// 方式二：自动类型推导
	var num = 10.11
	fmt.Println(num)

	// 方式三：省略var，:=左变量不应该是已经声明过的，否则会编译错误，且冒号不能省略
	// name = "abc"  典型错误，先定义，再赋值
	name := "abc" // 等价于 val name string      name = "abc"
	fmt.Println(name)

	// 多变量声明
	//
	// 多变量声明，类型相同
	/*
	 var a int
	 var b int
	 var c int
	*/
	var a, b, c int
	fmt.Println(a, b, c)

	// 多变量声明，类型不相同
	// 等价于 n1=100,n2="tom",n3=123
	var n1, n2, n3 = 100, "tom", 123
	fmt.Println(n1, n2, n3)

	// 多变量声明，类型推导
	a1, a2, a3 := 101, "jack", 321
	fmt.Println(a1, a2, a3)

	// 输出全局变量
	fmt.Println(b1, b2, b3)

	fmt.Println(n4, n5, n6)

	// fmt.Println(n7, n8, n9)
}
