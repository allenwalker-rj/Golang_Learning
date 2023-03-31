package main

import "fmt"

func main() {
	num := 10
	test(num)
	fmt.Println("test num =", num)
	fmt.Println("&num =", &num)
	test2(&num)
	fmt.Println("test2 num =", num)
}

// 基础数据类型 和 数组 默认都是值传递，即进行拷贝。在函数内修改，不会影响到原来的值
func test(n int) {
	n = 30
	fmt.Println("&n =", &n)
	fmt.Println("n =", n)
}

// 以值传递方式的数据类型，如果希望在函数内的变量能修改函数外的变量，可以传入变量的地址&
// 函数内内以指针的方式操作变量，从效果来看类型引用传递
func test2(n *int) {
	fmt.Println("&*n =", &*n)
	*n = 30
	fmt.Println("n =", *n)
}
