package main

import "fmt"

// 变量使用注意事项
func main() {
	// 变量的重新赋值
	// 不可以更改数据类型、不可在函数外修改局部变量
	var i int = 10
	i = 30
	i = 40

	fmt.Println(i)
	// i = 1.2		// 不可修改数据类型

	// 同作用域下变量不可重名
	// 一个大括号就是一个作用域
	{
		var i int = 100
		fmt.Println(i)
		{
			i := 200
			fmt.Println(i)
		}
	}
	// 变量三要素：变量名+值+数据类型
	// 如果没有赋初值，则编译器自动使用默认值，int为0，string为空串，小数也为0
}
