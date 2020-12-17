package main

import (
	"fmt"
)

// 标识符的使用
func main() {
	// 区分大小写
	var num int = 10
	var Num int = 20
	fmt.Println(num, Num)

	// 不能包含空格
	// syntax 语法错误
	// var ab cd int = 40 // syntax error: unexpected int at end of statementgo
	// fmt.Println(ab cd)

	// _ 仅能当做占位符，用于忽略某个值，来使用
	// 而不能直接当做标识符来使用
	// 但可以和字母、数字合法的组成标识符
	// var _ int = 30
	// fmt.Println(_) // cannot use _ as valuego

	// go中不建议使用 _ 命名方式，甚至不建议在标识符中出现下划线
	// don't use underscores in Go names; var a_b should be aB
	// var a_b int = 30
	// fmt.Println(a_b)

	// 不能和系统的 25 个保留关键字重名
	// 但是 int 和 float32 等都没在里面，故都可以命名为标识符
}
