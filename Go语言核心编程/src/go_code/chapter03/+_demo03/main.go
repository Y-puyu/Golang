package main

import "fmt"

// +的使用
func main() {
	var a, b int
	a = 1
	b = 1
	var c = a + b
	fmt.Println(c)

	var str1, str2 = "a", "b"
	var str3 = str1 + str2
	fmt.Println(str3)
}
