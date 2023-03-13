package main

import "fmt"

// 全局变量，定义在函数外的变量
var globe string = "这个是一个全局变量"

// 全局变量的另一种写法
var (
	aaa int     = 10
	bbb float32 = 11.11
	ccc float64 = 12.12
)

func main() {
	// 变量的声明
	var age int
	// 变量的赋值
	age = 18
	fmt.Println("age = ", age)

	// 声明和赋值一起
	var age2 int = 19
	fmt.Println("age2 = ", age2)

	// 指定变量的类型，但是不赋值，使用默认值
	var age3 int
	fmt.Println("age3 = ", age3)

	// 省略 var , 使用 := 赋值
	gender := "男"
	fmt.Println("gender = ", gender)

	fmt.Println("--------------------------------------------------------")

	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	var n4, name, n5 = 10, "jack", 7.8
	fmt.Println(n4, name, n5)

	fmt.Println(globe)
	fmt.Println(aaa, bbb, ccc)

}
