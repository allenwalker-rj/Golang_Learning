package main

import "fmt"

func main() {
	fmt.Println("service 10")
	a := test
	fmt.Printf("a的类型是: %T, test函数的类型是: %T", a, test)
	fmt.Println()
	// 通过该变量可以对函数进行调用
	a(10)

	test2(10, 3.14, test)
	test2(10, 3.14, a)

	// 为了简化数据类型定义，GO支持自定义数据类型
	// 基本语法：type 数据类型名 [类型别名]
	type myInt int
	var b myInt = 30
	fmt.Println("b =", b)

	// 虽然是别名，但是在 GO 种编译识别时还是会任务 myInt 和 int 不是同一种数据类型
	var c int = 40
	c = int(b)
	fmt.Println("c =", c)

	test3(a)

}

// 函数也是一种数据类型，可以赋值给一个变量
func test(num int) {
	fmt.Println("test num =", num)
}

// 定义一个函数，把另一个函数作为形参
func test2(num1 int, num2 float32, testFunc func(int)) {
	fmt.Println("-----test2")
}

type myFunc func(int)

func test3(testFunc myFunc) {
	fmt.Println("-----test3")
}
