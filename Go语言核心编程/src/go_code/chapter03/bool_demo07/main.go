package main

import (
	"fmt"
	"unsafe"
)

// bool类型的使用
// 只允许为true或false，不允许为0,1,非零值
// 占用1个字节
func main() {
	var flag bool = false
	fmt.Println(flag)                // false
	fmt.Println(unsafe.Sizeof(flag)) // 1
	// flag = 1 	// 错误使用
}
