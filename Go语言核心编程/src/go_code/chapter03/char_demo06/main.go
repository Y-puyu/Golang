package main

import (
	"fmt"
)

// 字符的使用
// 没有专门的字符类型，若存储单个字符，一般使用byte来保存
// Go的字符串由字节组成
// 英文一个字节，汉字三个字节
func main() {
	var c1 byte = 'a' // 0~255
	var c2 byte = 'b'
	fmt.Println(c1, c2, c1+c2) // 97 98 195，输出的是ascll码值

	fmt.Printf("%c", c1) // a		格式化输出

	// var c3 byte = '北'		// constant 21271 overflows byte，溢出，utf-8对应查Unicode值为21271
	var c3 int = '北' // 对应码值大于255，考虑使用int类型
	var c4 int = 21271
	fmt.Printf("%c", c3) // 北
	fmt.Printf("%c", c4) // 北

	// 字符的转义，在chapter02中已经学习了
}
