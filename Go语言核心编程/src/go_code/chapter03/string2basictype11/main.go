package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型和string的转换
func main() {
	var s1 string = "true"
	var n1 bool

	// 说明：ParseBool(str) 会返回两个值 (value bool, err error)
	// 只用 b 来接收第一个返回值 value bool，而使用 _ 来忽略第二个返回值
	n1, _ = strconv.ParseBool(s1)
	fmt.Printf("%v %T\n", n1, n1) // true bool

	var s2 string = "12345"
	var n2 int64
	n2, _ = strconv.ParseInt(s2, 10, 0) // ParseInt 返回值是 int64。s2 字符串，10 代表 10 进制，0 代表 int类型， 8代表 int8
	fmt.Printf("%v %T\n", n2, n2)       // 12345 int64

	var s3 string = "123.3435"
	var n3 float64
	n3, _ = strconv.ParseFloat(s3, 64) // ParseFloat  返回值是 float64，64 代表类型
	fmt.Printf("%v %T\n", n3, n3)      // 123.3435 float64

	// 返回值类型一定要按照类型接收，若想转换，可以自行强转。不能拿非对应类型来接收

	// 注意：
	// 基本类型转 string 一般不会出问题。
	// 但是基本类型转 string 可能会导致一系列问题
	// 若无法成功转换，则会给转换对象填充默认值

	var s4 string = "hello"
	var n4 int64 = 11
	n4, _ = strconv.ParseInt(s4, 10, 64) // 无法将 "hello" 转成 int64
	fmt.Printf("%v %T\n", n4, n4)        // 0 int64。会将原来的值转成默认的 0 值

	// 只要是非法转换，编译运行均不会报错。但是会将原来实体填充默认值
	// int、float变成 0 值，bool 变为 false。且会将原来的值覆盖
}
