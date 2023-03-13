package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// int8 范围 -128 ~ 127
	var a int8 = 127
	fmt.Println(a)

	// unit8 范围 0 ~ 255
	var b uint8 = 255
	fmt.Println(b)

	// rune 等价 int32
	var c rune = -1
	fmt.Println(c)

	// byte 等价 uint8
	var d byte = 0
	fmt.Println(d)

	// Printf 函数的作用就是格式化
	var e = 28
	// %T 输出对应值的类型
	fmt.Printf("e的类型是: %T", e)
	fmt.Println()
	fmt.Println("e拥有的字节数: ", unsafe.Sizeof(e))

}
