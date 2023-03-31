package main

import "fmt"

func main() {
	fmt.Println("service-8")
	test()
	test(1)
	test(1, 2)
	test(1, 2, 3)
	test(1, 2, 3, 4)
	test(1, 2, 3, 4, 5)
}

// 定义一个可变参数的函数， args ...int 可以传入任意多个数量的int类型的数据
func test(args ...int) {
	fmt.Println(args)
}
