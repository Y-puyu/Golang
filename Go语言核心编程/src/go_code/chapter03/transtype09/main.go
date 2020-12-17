package main

import "fmt" // 如果我们没有使用一个包，但不想删除。可以前面加一个下划线表示将其忽略

func main() {
	var i int32 = 100
	// var n1 float32 = i		// 经典错误，不能直接不同类型直接赋值
	var n1 float32 = float32(i) // 显示转换
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Println(n1, n2, n3) // 100 100 100

	// 不管是低精度到高精度，高精度到低精度，转换会不会发生精度丢失或者数字截断
	// 都必须显示强制转换。不支持隐式类型转换

	// 转换过后，i的数据类型并没有发生变化，只不过是将值交给了被转换对象
	fmt.Printf("%T %T %T %T \n", i, n1, n2, n3) // int32 float32 int8 int64

	// 如果将int64转换为int8，放不下的时候。
	// 编译不会报错，只不过转换的结果按溢出处理，与预期值不同
	var a1 int64 = 99999
	var a2 int8 = int8(a1) // 按溢出处理。在二进制中在详讲。在此仅简单了解即可
	fmt.Println(a1, a2)    // 99999 -97

	// 练习 1
	var b1 int32 = 12
	var b2 int64
	var b3 int8

	// b2 = b1 + 20 // n1 + 20 的数据类型是int32，不能将int32交给int64
	// b3 = b1 + 20 // 同理

	b2 = int64(b1 + 20) // 或者 int64(b1) + 20
	b3 = int8(b1 + 20)  // 或者 int8(b1) + 20

	fmt.Println(b2, b3) // 32 32

	var b4 int8
	b4 = int8(b1) + 127 // 编译通过，但是 12 + 127 运行时结果按照溢出处理
	// b3 = int8(b1) + 128 // 编译不过，128 已经越界了。报错：constant 128 overflows int8go
	fmt.Println(b4) // -117
}
