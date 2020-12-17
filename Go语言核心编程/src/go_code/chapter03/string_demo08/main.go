package main

import (
	"fmt"
)

// string类型的使用
func main() {
	var a string = "北京长城hello"
	fmt.Println(a) // 北京长城hello

	// 使用细节
	// 1.统一使用utf-8编码，对中文友好。不会产生乱码问题。

	// 2.字符串不可变
	var str = "hello"
	fmt.Println(str)
	// str[0] = 'a'

	// 3.字符串两种表示形式
	// 3.1 双引号，会识别转义字符
	str1 := "abc\ndef"
	fmt.Println(str1) // abc换行输出def
	// 3.2 反引号，以字符串的原生形式输出，包括换行和特殊字符，
	str2 := `
		// 3.字符串两种表示形式
	// 3.1 双引号，会识别转义字符
	str1 := "abc\ndef"
	fmt.Println(str1)	// abc换行输出def
	// 3.2 反引号，以字符串的原生形式输出，包括换行和特殊字符，
	`
	fmt.Println(str2) // 整体输出，且保留原格式。思考字符串中包含反引号怎么办？

	// 4.字符串拼接方式
	str3 := "he" + "oll"
	str3 += "   world~"
	fmt.Println(str3) // heoll   world~

	// 5.针对长字符串拼接
	// 可以采用换行操作,但一定需要将+保留到行尾，否则go自己加分号就玩完了
	str4 := "he" + "oll" + "he" + "oll" + "he" + "oll" + "he" + "oll" +
		"he" + "oll" + "he" + "oll" + "he" + "oll" +
		"he" + "oll"
	fmt.Println(str4) // heollheollheollheollheollheollheollheoll

	// go各种基本类型的默认值，也称为0值
	var g int     // 0
	var b float32 // 0
	var c float64 // 0
	var d bool    //false
	var e string  // ""
	// %v 按默认格式输出
	fmt.Printf("%d %f %f %v %v\n", g, b, c, d, e) // 0 0.000000 0.000000 false
	fmt.Printf("%d %v %v %v %v", g, b, c, d, e)   // 0 0 0 false
}
