package main

import "fmt"

func main() {
	//定义浮点类型数据
	var num1 float32 = 3.14
	fmt.Println(num1)
	var num2 float32 = -3.14
	fmt.Println(num2)
	var num3 float32 = 314e-2
	fmt.Println(num3)
	var num4 float32 = 314e+2
	fmt.Println(num4)
	var num5 float32 = 314e+2
	fmt.Println(num5)
	var num6 float64 = 314e+2
	fmt.Println(num6)

	// 浮点数，可能会有精度的损失，通常情况下建议使用 float64
	var num7 float32 = 256.000000916
	fmt.Println(num7)
	var num8 float64 = 256.000000916
	fmt.Println(num8)

	// golang 种默认的浮点类型为：float64
	var num9 = 3.14
	fmt.Printf("num9的类型是: %T", num9)
}
